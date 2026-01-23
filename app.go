package main

import (
	"Clipcat/internal/clipboard"
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	gclip "golang.design/x/clipboard"
)

// to read clipboard contents, use cli command <<< go get github.com/atotto/clipboard >>>

// App struct
type App struct {
	ctx context.Context
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
		// Try image first\

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

// GetClips returns all clips from the database
func (a *App) GetClips() ([]Clip, error) {
	return getClips()
}

// TogglePin toggles the pinned status of a clip
func (a *App) TogglePin(clipID int) error {
	return togglePinClip(clipID)
}

// Delete a clip by ID
func (a *App) Delete(clipID int) error {
	return deleteClip(clipID)
}

// GetStorageLimit returns the current storage limit
func (a *App) GetStorageLimit() (int, error) {
	return getStorageLimit()
}

// UpdateStorageLimit updates the storage limit
func (a *App) UpdateStorageLimit(newLimit int) error {
	return updateStorageLimit(newLimit)
}

// get app data directory
func getAppDataDir() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	appDir := filepath.Join(dir, "clipussy/db")
	err = os.MkdirAll(appDir, 0755)

	return appDir, err
}
