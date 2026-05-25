//go:build windows

package platform

import (
	"syscall"
	"unsafe"
)

var (
	modKernel32             = syscall.NewLazyDLL("kernel32.dll")
	modUser32               = syscall.NewLazyDLL("user32.dll")
	procCreateMutexW        = modKernel32.NewProc("CreateMutexW")
	procFindWindowW         = modUser32.NewProc("FindWindowW")
	procShowWindow          = modUser32.NewProc("ShowWindow")
	procSetForegroundWindow = modUser32.NewProc("SetForegroundWindow")
	procIsIconic            = modUser32.NewProc("IsIconic")
)

const (
	errAlreadyExists = syscall.Errno(183)
	swRestore        = 9
	swShow           = 5
)

// EnsureSingleInstance returns true if this is the only running instance.
// If another instance is already running it brings that window to the
// foreground and returns false, so main() can exit immediately.
func EnsureSingleInstance() bool {
	mutexName, _ := syscall.UTF16PtrFromString("Local\\ClipCatSingleInstanceMutex_v1")
	_, _, lastErr := procCreateMutexW.Call(
		0,
		0,
		uintptr(unsafe.Pointer(mutexName)),
	)
	if lastErr != errAlreadyExists {
		return true
	}

	title, _ := syscall.UTF16PtrFromString("Clipcat")
	hwnd, _, _ := procFindWindowW.Call(0, uintptr(unsafe.Pointer(title)))
	if hwnd != 0 {
		iconic, _, _ := procIsIconic.Call(hwnd)
		if iconic != 0 {
			procShowWindow.Call(hwnd, swRestore)
		} else {
			procShowWindow.Call(hwnd, swShow)
		}
		procSetForegroundWindow.Call(hwnd)
	}

	return false
}
