package main

import (
	_ "embed"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed build/windows/icon.ico
var trayIcon []byte

// startTray starts the system-tray icon in a background goroutine.
// It shows a tooltip of "Clipcat" and provides two menu items:
//   - Show Clipcat  → makes the window visible again
//   - Quit          → exits the application
//
// This is the only way to reach Clipcat after it hides itself following a paste.
func (a *App) startTray() {
	go systray.Run(a.onTrayReady, func() {})
}

func (a *App) onTrayReady() {
	systray.SetIcon(trayIcon)
	systray.SetTitle("Clipcat")
	systray.SetTooltip("Clipcat – press Ctrl+Shift+V to open")

	mShow := systray.AddMenuItem("Show Clipcat", "Bring the Clipcat window to the front")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Exit Clipcat completely")

	// Handle menu clicks.
	go func() {
		for {
			select {
			case <-mShow.ClickedCh:
				if a.ctx != nil {
					runtime.WindowShow(a.ctx)
					runtime.WindowSetAlwaysOnTop(a.ctx, true)
					runtime.WindowSetAlwaysOnTop(a.ctx, false)
				}
			case <-mQuit.ClickedCh:
				systray.Quit()
				if a.ctx != nil {
					runtime.Quit(a.ctx)
				}
			}
		}
	}()
}
