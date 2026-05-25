//go:build darwin

package clipboard

import (
	"os/exec"
	"strings"
	"sync"
	"time"
)

//
// Previous app tracking (macOS)
//

var (
	prevAppBundleID string
	prevAppMu       sync.Mutex
	ourPID          uint32
)

// SetOurProcessID stores this process's PID (not heavily used on macOS, but
// kept for API compatibility).
func SetOurProcessID(pid uint32) {
	ourPID = pid
}

// StartFocusTracker polls the frontmost application every 150 ms and stores
// its bundle ID. The polled value is used by FocusPreviousWindow.
func StartFocusTracker() {
	go func() {
		for {
			time.Sleep(150 * time.Millisecond)
			name := getForegroundAppNameDarwin()
			if name != "" {
				prevAppMu.Lock()
				prevAppBundleID = name
				prevAppMu.Unlock()
			}
		}
	}()
}

// capturePreviousAppDarwin snapshots the current frontmost app bundle ID.
func capturePreviousAppDarwin() {
	name := getForegroundAppNameDarwin()
	if name != "" {
		prevAppMu.Lock()
		prevAppBundleID = name
		prevAppMu.Unlock()
	}
}

// HasPreviousWindow reports whether a non-Clipcat app has been tracked.
func HasPreviousWindow() bool {
	prevAppMu.Lock()
	defer prevAppMu.Unlock()
	return prevAppBundleID != ""
}

// FocusPreviousWindow activates the previously tracked application via osascript.
func FocusPreviousWindow() {
	prevAppMu.Lock()
	id := prevAppBundleID
	prevAppMu.Unlock()

	if id == "" {
		return
	}
	exec.Command("osascript", "-e",
		`tell application id "`+id+`" to activate`,
	).Run()
}

// SimulatePaste sends a Cmd+V keystroke via osascript.
func SimulatePaste() {
	time.Sleep(80 * time.Millisecond)
	exec.Command("osascript", "-e",
		`tell application "System Events" to keystroke "v" using command down`,
	).Run()
}

//
// Process ignore list – macOS implementation
//

// isForegroundProcessIgnored returns true if the frontmost app's bundle ID
// matches any entry in the ignore list.
func isForegroundProcessIgnored() bool {
	ignoredProcessesMu.RLock()
	defer ignoredProcessesMu.RUnlock()

	if len(ignoredProcesses) == 0 {
		return false
	}

	name := getForegroundAppNameDarwin()
	if name == "" {
		return false
	}

	for _, ignored := range ignoredProcesses {
		if strings.Contains(name, ignored) {
			return true
		}
	}
	return false
}
