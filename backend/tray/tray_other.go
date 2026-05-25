//go:build !windows

package tray

// Start is a no-op on non-Windows platforms.
// The app is accessible via the macOS dock / Linux taskbar.
func Start(iconBytes []byte, onShow func(), onQuit func()) {
	// No system tray on non-Windows platforms for now.
	// The tray icon can be added later with a custom CGo implementation.
}
