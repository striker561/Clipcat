package clipboard

import (
	"sync"
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"

	"github.com/lxn/win"
)

const (
	WM_CLIPBOARDUPDATE = 0x031D
	WM_HOTKEY          = 0x0312

	// Global hotkey: Ctrl+Shift+V (ID 1)
	hotkeyID    = 1
	MOD_SHIFT   = 0x0004
	MOD_CONTROL = 0x0002
	VK_V        = 0x56
)

var (
	lastClipboardChange time.Time
	clipboardMutex      sync.Mutex

	onChangeCallback func()
	onHotkeyCallback func()

	// isPaused stops new clips from being saved without stopping the listener.
	isPaused atomic.Bool
)

// PauseCapture stops new clipboard events from being saved.
func PauseCapture() { isPaused.Store(true) }

// ResumeCapture re-enables clipboard capture.
func ResumeCapture() { isPaused.Store(false) }

// IsPaused reports whether capture is currently paused.
func IsPaused() bool { return isPaused.Load() }

// wndProc handles messages for the hidden clipboard + hotkey window.
// Must be a package-level function so the GC never collects the callback pointer.
func wndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {

	case WM_CLIPBOARDUPDATE:
		if isPaused.Load() {
			return 0
		}
		if isForegroundProcessIgnored() {
			return 0
		}

		now := time.Now()
		clipboardMutex.Lock()
		if now.Sub(lastClipboardChange) > 150*time.Millisecond {
			lastClipboardChange = now
			clipboardMutex.Unlock()
			if onChangeCallback != nil {
				onChangeCallback()
			}
		} else {
			clipboardMutex.Unlock()
		}
		return 0

	case WM_HOTKEY:
		if wParam == hotkeyID {
			// Capture the foreground window now, before Clipcat steals focus.
			capturePreviousWindow()
			if onHotkeyCallback != nil {
				go onHotkeyCallback()
			}
		}
		return 0
	}

	return win.DefWindowProc(hwnd, msg, wParam, lParam)
}

// StartClipboardListener creates a hidden message-only window that:
//   - Listens for clipboard changes, calling onChange (respects pause + ignore list)
//   - Listens for the Ctrl+Shift+V global hotkey, calling onHotkey
func StartClipboardListener(onChange func(), onHotkey func()) {
	onChangeCallback = onChange
	onHotkeyCallback = onHotkey

	go func() {
		instance := win.GetModuleHandle(nil)
		className, _ := syscall.UTF16PtrFromString("ClipcatClipboardWindow")

		var wc win.WNDCLASSEX
		wc.CbSize = uint32(unsafe.Sizeof(wc))
		wc.LpfnWndProc = syscall.NewCallback(wndProc)
		wc.HInstance = instance
		wc.LpszClassName = className

		if win.RegisterClassEx(&wc) == 0 {
			panic("clipboard: failed to register window class")
		}

		hwnd := win.CreateWindowEx(
			0, className, nil, 0,
			0, 0, 0, 0,
			win.HWND_MESSAGE, // hidden message-only window
			0, instance, nil,
		)
		if hwnd == 0 {
			panic("clipboard: failed to create message window")
		}

		if !win.AddClipboardFormatListener(hwnd) {
			panic("clipboard: failed to register clipboard format listener")
		}

		// Ctrl+Shift+V  global hotkey to show/hide Clipcat
		procRegisterHotKey.Call(uintptr(hwnd), hotkeyID, MOD_CONTROL|MOD_SHIFT, VK_V)

		var msg win.MSG
		for win.GetMessage(&msg, 0, 0, 0) > 0 {
			win.TranslateMessage(&msg)
			win.DispatchMessage(&msg)
		}
	}()
}
