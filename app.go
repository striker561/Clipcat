package main

import (
	"Clipcat/internal/clipboard"
	"Clipcat/internal/startup"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	gclip "golang.design/x/clipboard"
)

// App struct
type App struct {
	ctx        context.Context
	isMiniClip bool
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// exposes app version to frontend
func (a *App) GetVersion() string {
	return AppVersion
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// initialize clipboard
	err := gclip.Init()
	if err != nil {
		panic(err)
	}

	appDir, err := getAppDataDir()
	if err != nil {
		panic(err)
	}

	dbPath := filepath.Join(appDir, "gyatt.db")

	err = InitDB(dbPath)
	if err != nil {
		panic(err)
	}

	createTables()
	migrateClipsTable()

	// Sync the ignore list from the DB into the in-memory clipboard filter.
	if ignoreList, err := getIgnoreList(); err == nil {
		clipboard.SetIgnoredProcesses(ignoreList)
	}

	// start clipboard listener
	var lastImage []byte
	clipboard.StartClipboardListener(func() {
		// Try image first

		if img := gclip.Read(gclip.FmtImage); img != nil {
			lastImagePtr := &lastImage

			if string(*lastImagePtr) == string(img) {
				// same image as before, skip
				return
			}

			// new image, save it
			*lastImagePtr = img
			err := addImageClip(img)
			if err != nil {
				fmt.Println("failed to save image:", err)
			}
			if a.ctx != nil {
				runtime.EventsEmit(a.ctx, "clipboard:changed")
			}
			return
		}

		// Fallback to text
		text := string(gclip.Read(gclip.FmtText))
		if text == "" {
			return
		}

		err := addClip(text, "text")
		if err != nil {
			fmt.Println("failed to save text:", err)
			return
		}

		if a.ctx != nil {
			runtime.EventsEmit(a.ctx, "clipboard:changed", text)
		}
	}, func() {
		// Hotkey (Ctrl+Shift+V) fired — show Clipcat and bring it to the front.
		if a.ctx == nil {
			return
		}
		runtime.WindowShow(a.ctx)
		runtime.WindowSetAlwaysOnTop(a.ctx, true)
		time.Sleep(150 * time.Millisecond)
		runtime.WindowSetAlwaysOnTop(a.ctx, false)
	})
}

func getAppDataDir() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	appDir := filepath.Join(dir, "clipussy/db")
	err = os.MkdirAll(appDir, 0755)

	return appDir, err
}

// --------------------------------------------------------------------------------
// Storage Limit Functions
// --------------------------------------------------------------------------------
func (a *App) GetStorageLimit() (int, error) {
	return getStorageLimit()
}

func (a *App) UpdateStorageLimit(newLimit int) error {
	return updateStorageLimit(newLimit)
}

// --------------------------------------------------------------------------------
// Clip Management Functions
// --------------------------------------------------------------------------------
func (a *App) GetClips() ([]Clip, error) {
	return getClips()
}

func (a *App) UpdateClipContent(clipID int, newContent string) error {
	return updateClipContent(clipID, newContent)
}

func (a *App) TogglePin(clipID int) error {
	return togglePinClip(clipID)
}

func (a *App) Delete(clipID int) error {
	return deleteClip(clipID)
}
func (a *App) DeleteAllClips() error {
	return deleteAllClips(a.ctx)
}

func (a *App) DeletePinnedClips() error {
	return deletePinnedClips(a.ctx)
}

func (a *App) DeleteUnpinnedClips() error {
	return deleteUnpinnedClips(a.ctx)
}

// --------------------------------------------------------------------------------
// Capture Pause / Resume
// --------------------------------------------------------------------------------

func (a *App) PauseCapture() {
	clipboard.PauseCapture()
}

func (a *App) ResumeCapture() {
	clipboard.ResumeCapture()
}

func (a *App) IsPaused() bool {
	return clipboard.IsPaused()
}

// --------------------------------------------------------------------------------
// Ignore List
// --------------------------------------------------------------------------------

func (a *App) GetIgnoreList() ([]string, error) {
	return getIgnoreList()
}

func (a *App) AddIgnoreEntry(name string) error {
	if err := addIgnoreEntry(name); err != nil {
		return err
	}
	// Keep the in-memory filter in sync.
	if list, err := getIgnoreList(); err == nil {
		clipboard.SetIgnoredProcesses(list)
	}
	return nil
}

func (a *App) RemoveIgnoreEntry(name string) error {
	if err := removeIgnoreEntry(name); err != nil {
		return err
	}
	if list, err := getIgnoreList(); err == nil {
		clipboard.SetIgnoredProcesses(list)
	}
	return nil
}

// --------------------------------------------------------------------------------
// Paste to Previous Window
// --------------------------------------------------------------------------------
//
// Sets the clipboard to the given text, hides Clipcat, re-focuses the window
// that was active when the hotkey was pressed, then simulates Ctrl+V.

func (a *App) PasteToWindow(content string) error {
	// Write content to the system clipboard.
	gclip.Write(gclip.FmtText, []byte(content))

	// Hide Clipcat so it gets out of the way before we paste.
	runtime.WindowHide(a.ctx)
	time.Sleep(120 * time.Millisecond)

	// Restore focus to where the user was, then fire Ctrl+V.
	clipboard.FocusPreviousWindow()
	clipboard.SimulatePaste()
	return nil
}

func (a *App) AddClip(content string, pinned bool) error {
	err := addManualClip(content, pinned)
	if err != nil {
		return err
	}
	if a.ctx != nil {
		runtime.EventsEmit(a.ctx, "clipboard:changed")
	}
	return nil
}

// --------------------------------------------------------------------------------
// Mini Clip Mode Functions
// --------------------------------------------------------------------------------
func (a *App) makeMiniClip(value bool) {
	if a.isMiniClip == value {
		return
	}

	runtime.WindowUnmaximise(a.ctx)
	runtime.WindowSetAlwaysOnTop(a.ctx, value)

	if value {
		runtime.WindowSetPosition(a.ctx, 20, 20)
		runtime.WindowSetMaxSize(a.ctx, 500, 300)
	} else {
		runtime.WindowSetMaxSize(a.ctx, 0, 0)
	}

	a.isMiniClip = value
}

func (a *App) MakeMiniClip(value bool) {
	a.makeMiniClip(value)
}

func (a *App) IsMiniClip() bool {
	return a.isMiniClip
}

//
// --------------------------------------------------------------------------------
// Windows Startup Management Functions
// --------------------------------------------------------------------------------
//

func (a *App) EnableStartup() error {
	return startup.EnableStartupWindows()
}

func (a *App) DisableStartup() error {
	return startup.RemoveStartupWindows()
}

func (a *App) IsStartupEnabled() (bool, error) {
	return startup.IsStartupEnabledWindows()
}
