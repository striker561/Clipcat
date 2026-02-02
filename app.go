package main

import (
	"Clipcat/internal/clipboard"
	"Clipcat/internal/startup"
	"context"
	"fmt"
	"os"
	"path/filepath"

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

//
// --------------------------------------------------------------------------------
// Storage Limit Functions
// --------------------------------------------------------------------------------
//
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
