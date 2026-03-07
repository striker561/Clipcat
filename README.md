<p align="center">
  <img src="./build/appicon.png" alt="icon" width="90">
</p>
<div align="center">
<h1>Clipcat</h1>
</div>

A creative and stylish clipboard manager built with Wails, designed to keep track of everything you copy through a clean, paper-aesthetic interface. It automatically records every clipboard change in real time, storing your copy history so you can easily revisit, reuse, and manage past content whenever you need it.

<img width="1912" height="1026" alt="image" src="https://github.com/user-attachments/assets/ec4870c1-8555-49e0-9cb4-52c83f7551b0" />


## Download

![Clipcat Banner](https://img.shields.io/badge/Made%20with-Wails-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/license-MIT-blue?style=for-the-badge)

[![Download Clipcat v0.8.0 for Windows](https://img.shields.io/badge/Download-Windows%20Installer-0078D4?style=for-the-badge&logo=windows&logoColor=white)](https://github.com/d3uceY/Clipcat/releases/download/v0.8.0/Clipcat.exe)

**[Download Clipcat](https://github.com/d3uceY/Clipcat/releases/download/v0.8.0/Clipcat.exe)**

> Windows 10/11 (64-bit) | Version 0.8.0

![clipcat (1)](https://github.com/user-attachments/assets/ca28ae42-2a9d-42c3-9a47-183808d59cf6)

### Windows SmartScreen Warning

When running the app for the first time, Windows SmartScreen may show a warning because the app is not yet code-signed. This is normal for open-source applications.

**To run the app:**
1. Click "More info" on the SmartScreen warning
2. Click "Run anyway"

Or alternatively:
1. Right-click the downloaded .exe file
2. Select "Properties"
3. Check "Unblock" at the bottom
4. Click "Apply" then "OK"
5. Run the executable

The app is safe and [open source](https://github.com/d3uceY/Clipcat) - you can verify the code yourself!


## Features

- **Automatic Clipboard Monitoring** - Automatically captures everything you copy (text and images) in real time via a native Windows message-only window

- **Image Support** - Captures and displays clipboard images, stored as BLOBs and rendered from base64 in the UI

- **Pin Important Clips** - Keep your most-used clips anchored to the top; pinned clips are protected from automatic deletion when the storage limit is reached
- <img width="577" height="471" alt="image" src="https://github.com/user-attachments/assets/aa6e7ff1-159e-4fc0-8565-fdc9208a07c4" />

- **Fast Search** - Instantly filter all clips with `Ctrl+F`; searches both pinned and recent clips simultaneously
- <img width="530" height="112" alt="showcase-search" src="https://github.com/user-attachments/assets/c2c76d50-5f94-481d-9191-ad37f2518967" />

- **Unique Paper Aesthetic** - Beautiful hand-drawn, notebook-style UI with GSAP animations, paper curtain reveals, and a soundtrack of satisfying sounds

- **Easy Management** - Copy, pin, and delete clips with intuitive hover controls

- **Edit Clips** - Modify the text content of any saved clip in-place without re-copying

- **Manual Clip Creation** - Add new clips directly from the app without copying anything; supports pinning on creation
- <img width="705" height="324" alt="showcase-action-btns" src="https://github.com/user-attachments/assets/d03a6634-8b41-4d78-a976-662b4c2b8f89" />

- **Privacy Mode** - Instantly blur all clip content for privacy or screen sharing; toggle with `Alt+H`
- <img width="242" height="227" alt="image" src="https://github.com/user-attachments/assets/9fed26d4-31cb-4ff0-bdd4-e8ee25780dba" />
- <img width="1223" height="244" alt="image" src="https://github.com/user-attachments/assets/0d242347-2e1d-46bf-b57f-20255d7c4fd1" />

- **Quick Paste** - A power-user workflow mode. When enabled, Clipcat hides to the system tray. Press `Ctrl+Shift+V` from any window to summon it, click the paste icon on a clip, and it fires directly into the window you were just using -- then vanishes again. When disabled, the paste button still lets you fire any clip into your last focused window without hiding the app.
- <img width="331" height="283" alt="image" src="https://github.com/user-attachments/assets/65dcb9dd-a402-4699-97fa-3042f1a4a9aa" />


- **System Tray** - Clipcat lives in the system tray and can be summoned or quit from there at any time, even when the window is hidden
- <img width="211" height="108" alt="image" src="https://github.com/user-attachments/assets/078cf695-5af4-41f5-a357-2ac21d12fc95" />


- **Global Hotkey** - `Ctrl+Shift+V` is a system-wide hotkey that brings Clipcat to the front from any application

- **Paste Into Any Window** - Every text clip has a paste button that fires its content directly into whichever window you had focused before opening Clipcat. No manual copying needed -- the previous window is tracked automatically in the background

- **Blocked Apps** - Add any `.exe` name (e.g. `1password.exe`) to a blocklist. Clipboard changes originating from those processes will be silently ignored and never stored
- <img width="259" height="335" alt="image" src="https://github.com/user-attachments/assets/ca3ed1ef-8a44-4aa2-a759-37d025e0682b" />

- **Pause Capture** - Temporarily suspend all clipboard monitoring without closing the app. Resume it any time from settings
- <img width="252" height="233" alt="image" src="https://github.com/user-attachments/assets/2ce57fca-f3bc-49c1-bfde-0ef62338dcde" />


- **Bulk Actions** - Quickly delete all clips, only pinned clips, or only unpinned clips via the settings menu -- all with a confirmation dialog
- <img width="196" height="179" alt="image" src="https://github.com/user-attachments/assets/a6cb09b3-14fa-4bc7-884a-57fc0d17d561" />
- <img width="467" height="193" alt="image" src="https://github.com/user-attachments/assets/a78a8a32-0d91-4920-ba4a-4100bb8d8cca" />

- **Hyperlink Detection** - URLs inside clip text are automatically rendered as clickable links that open in your default browser

- **Relative Timestamps** - Each clip shows a live-updating relative time (e.g. "2 minutes ago", "yesterday")

- **Duplicate Detection** - Automatically prevents saving duplicate text or image content

- **Sound Effects** - Audible feedback for copy, paste, delete, pin, and settings interactions; can be toggled off with `Alt+S`

- **Persistent Storage** - SQLite database keeps your clips safe across restarts

- **Configurable Storage Limit** - Customize how many clips to keep (100-500, in steps of 50); pinned clips are always preserved regardless of the limit
- <img width="259" height="335" alt="image" src="https://github.com/user-attachments/assets/ca3ed1ef-8a44-4aa2-a759-37d025e0682b" />

- **Mini Clip Mode** - A compact, always-on-top window for unobtrusive usage; toggle with `Alt+M`
- <img width="270" height="330" alt="image" src="https://github.com/user-attachments/assets/ddedff44-5007-4f57-b98e-006e28e24e71" />

- **Startup Support** - Option to launch Clipcat automatically when your system starts
- <img width="293" height="335" alt="image" src="https://github.com/user-attachments/assets/2437742c-48c4-4e09-9726-19551d86eb54" />

- **Auto Update Check** - Automatically checks for new versions on GitHub


## Technologies Used

### Backend
- **[Go](https://golang.org/)** - Core application logic
- **[Wails v2](https://wails.io/)** - Desktop application framework
- **[SQLite](https://www.sqlite.org/)** (via modernc.org/sqlite) - Local database for clip storage
- **[golang.design/x/clipboard](https://github.com/golang-design/clipboard)** - Cross-platform clipboard access with image support
- **[lxn/win](https://github.com/lxn/win)** - Windows API bindings for the clipboard message window
- **[getlantern/systray](https://github.com/getlantern/systray)** - Cross-platform system tray icon and menu

### Frontend
- **[React 18](https://react.dev/)** - UI framework
- **[TypeScript](https://www.typescriptlang.org/)** - Type-safe JavaScript
- **[Vite](https://vitejs.dev/)** - Fast build tool and dev server
- **[Tailwind CSS](https://tailwindcss.com/)** - Utility-first CSS framework
- **[shadcn/ui](https://ui.shadcn.com/)** - Accessible component library
- **[GSAP](https://greensock.com/gsap/)** - Animation library
- **[Lucide React](https://lucide.dev/)** - Icon library

## Why Wails?

Clipcat uses Wails to:
- Integrate native Windows clipboard APIs via Go
- Communicate clipboard events to a React UI in real time
- Bundle a lightweight, native-feeling desktop app without Electron


## Architecture

Clipcat follows a clean architecture pattern with clear separation between frontend and backend:

```
+--------------------------------------------------+
|               Frontend (React)                   |
|  +------------+  +----------+  +-----------+    |
|  |  UI Layer  |  | Context  |  | Components|    |
|  | (TSX/CSS)  |  | Provider |  |  (Cards)  |    |
|  +------------+  +----------+  +-----------+    |
+---------------------+----------------------------+
                      | Wails Bridge (IPC)
+---------------------+----------------------------+
|              Backend (Go)                        |
|  +----------+  +----------+  +--------------+   |
|  |  app.go  |  | clips.go |  | settings.go  |   |
|  | (Bridge) |  | (Logic)  |  |   (Prefs)    |   |
|  +----------+  +----------+  +--------------+   |
|  +----------+  +----------+  +--------------+   |
|  |  tray.go |  | ignore.go|  |    db.go     |   |
|  |  (Tray)  |  | (Filter) |  |  (Storage)   |   |
|  +----------+  +----------+  +--------------+   |
+---------------------+----------------------------+
                      |
          +-----------+-----------+
          |                       |
    +-----+------+       +--------+--------+
    |  SQLite    |       |   Windows API   |
    | Database   |       | (Clipboard +    |
    +------------+       |  Hotkey + Focus)|
                         +-----------------+
```

### Data Flow

1. **Clipboard Monitoring**
   - A hidden `HWND_MESSAGE` window receives `WM_CLIPBOARDUPDATE` notifications from Windows
   - A 150 ms debounce prevents duplicate saves when apps write the clipboard in multiple steps
   - The ignore list filters out events from blocked processes before saving

2. **Focus Tracking**
   - A background goroutine polls `GetForegroundWindow` every 150 ms
   - It continuously stores the last non-Clipcat window, so the paste button always has a valid target without requiring the hotkey to be pressed first

3. **Data Storage**
   - New clips are saved to SQLite with duplicate detection (text and image)
   - Automatic cleanup keeps only the most recent N clips, always preserving pinned ones
   - App preferences (Quick Paste mode, storage limit, ignore list) are stored in the same database

4. **Frontend Updates**
   - Backend emits `clipboard:changed` events when new content is saved
   - React context manages all clip state and settings
   - UI automatically re-renders with new data

5. **User Actions**
   - Copy: Uses browser clipboard API
   - Paste to window: Focuses the previously tracked window and simulates Ctrl+V
   - Pin/Unpin: Toggles database flag, reorders UI
   - Delete: Removes from database, refreshes list
   - Search: Client-side filtering with instant results


## Project Structure

```
Clipcat/
|-- app.go                          # Wails app struct, startup, exposed bindings
|-- clips.go                        # Clip CRUD operations
|-- db.go                           # Database initialization and migrations
|-- settings.go                     # App settings (Quick Paste mode)
|-- ignore.go                       # Blocked app list persistence
|-- tray.go                         # System tray icon and menu
|-- main.go                         # Wails runtime setup
|-- go.mod                          # Go dependencies
|-- wails.json                      # Wails configuration
|-- internal/
|   `-- clipboard/
|       |-- listener_window.go      # Windows clipboard + hotkey listener
|       `-- window_utils.go         # Focus tracker, paste simulation, ignore check
|-- frontend/
|   |-- src/
|   |   |-- App.tsx                 # Root component
|   |   |-- components/
|   |   |   |-- page.tsx            # Main page layout
|   |   |   |-- clip-card.tsx       # Individual clip card with all actions
|   |   |   |-- window-controls.tsx # Title bar, settings panel
|   |   |   |-- add-clip-dialog.tsx # Manual clip creation dialog
|   |   |   |-- edit-clip-dialog.tsx# Clip edit dialog
|   |   |   |-- delete-clips-dialog.tsx # Bulk delete dialog
|   |   |   |-- about-dialog.tsx    # About / version dialog
|   |   |   `-- ui/                 # shadcn/ui base components
|   |   |-- context/
|   |   |   `-- ClipContext.tsx     # Global state management
|   |   |-- helpers/
|   |   |   |-- formatTime.ts       # Date formatting
|   |   |   |-- playSound.ts        # Audio feedback
|   |   |   |-- insertLinks.ts      # URL to clickable link renderer
|   |   |   |-- copyBase64Image.ts  # Image clipboard helper
|   |   |   `-- wait.ts             # Promise-based delay
|   |   `-- types/
|   |       `-- clip.ts             # TypeScript interfaces
|   |-- wailsjs/                    # Auto-generated Wails bindings
|   |-- public/                     # Static assets (sounds, images, cursors)
|   |-- package.json
|   `-- vite.config.ts
`-- build/                          # Build configuration
    `-- windows/
        `-- installer/              # NSIS installer config
```


## How It Works

### Backend Implementation

**Clipboard Monitoring** (`internal/clipboard/listener_window.go`)

A hidden `HWND_MESSAGE` window is registered with `AddClipboardFormatListener`. Windows delivers `WM_CLIPBOARDUPDATE` messages to it whenever any app writes to the clipboard. The same window also handles `WM_HOTKEY` for the global `Ctrl+Shift+V` shortcut via `RegisterHotKey`.

```go
case WM_CLIPBOARDUPDATE:
    // 150 ms debounce + pause check + ignore list check
    // then calls onChangeCallback()

case WM_HOTKEY:
    // Snapshot the current foreground window, then show Clipcat
    capturePreviousWindow()
    go onHotkeyCallback()
```

**Focus Tracking** (`internal/clipboard/window_utils.go`)

A background goroutine polls `GetForegroundWindow` every 150 ms, skipping Clipcat's own process, so the paste target is always up to date:

```go
func StartFocusTracker() {
    go func() {
        for {
            time.Sleep(150 * time.Millisecond)
            hwnd, _, _ := procGetForegroundWindow.Call()
            // skip our own PID, then store prevHWND
        }
    }()
}
```

**Database Schema** (`db.go`)

```sql
CREATE TABLE clips (
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    content    TEXT,
    image      BLOB,
    type       TEXT NOT NULL,
    pinned     BOOLEAN DEFAULT 0,
    created_at DATETIME
);

CREATE TABLE clip_storage_limit (
    id          INTEGER PRIMARY KEY CHECK (id = 0),
    limit_count INTEGER DEFAULT 100
);

CREATE TABLE ignore_list (
    process_name TEXT PRIMARY KEY
);

CREATE TABLE settings (
    id         INTEGER PRIMARY KEY CHECK (id = 0),
    ghost_mode INTEGER DEFAULT 0   -- 1 = Quick Paste enabled
);
```

> `content` is nullable to support image-only clips. Either `content` or `image` is populated based on the clip `type`.

**Key Operations** (`clips.go`)
- `getClips()` - Fetches all clips ordered by pinned status, then by date
- `clipExists()` / `imageClipExists()` - Prevent duplicate text and image clips
- `addClip()` - Inserts new text clip and enforces storage limit (preserving pinned)
- `addImageClip()` - Inserts new image clip as BLOB and enforces storage limit
- `addManualClip()` - Inserts a user-created clip with optional pinned flag
- `togglePinClip()` - Toggles pinned status by ID
- `deleteClip()` - Removes clip by ID
- `deleteAllClips()` / `deletePinnedClips()` / `deleteUnpinnedClips()` - Bulk delete with native confirmation dialog
- `updateClipContent()` - Edits text content of an existing clip

### Frontend Implementation

**State Management** (`ClipContext.tsx`)
- Global state with React Context
- Splits clips into `pinned` and `recent` arrays
- Listens for `clipboard:changed` events from backend
- Manages all settings: sound, hide content, mini clip, startup, pause capture, blocked apps, Quick Paste

**UI Components**
- **ClipCard** - Individual clip with copy / paste-to-window / edit / pin / delete actions; renders text (with clickable links) or image; shows relative timestamps
- **Page** - Main layout with animated paper curtain reveal, search bar, pinned section, and recent section
- **WindowControls** - Frameless title bar with minimize/maximize/close and an animated settings panel
- **AddClipDialog** - Inline manual clip creation with optional pinning
- **EditClipDialog** - Edit existing clip content in a modal
- **DeleteClipsDialog** - Bulk delete options (all / pinned / unpinned)
- **AboutDialog** - App info with automatic GitHub update checking

**Animations** (GSAP)
- Paper curtain reveal and cat character entrance on startup
- Settings panel slide/scale in and out
- Clip card row-span masonry layout
- Privacy mode cat swap
- Sound effects on every interaction


## Getting Started

### Prerequisites
- Go 1.24.0 or higher
- Node.js 18+ and npm
- Wails CLI: `go install github.com/wailsapp/wails/v2/cmd/wails@latest`

### Development

1. **Clone the repository**
   ```bash
   git clone https://github.com/d3uceY/clipcat.git
   cd clipcat
   ```

2. **Install dependencies**
   ```bash
   # Backend dependencies
   go mod download

   # Frontend dependencies
   cd frontend
   npm install
   cd ..
   ```

3. **Run in development mode**
   ```bash
   wails dev
   ```

   The app will launch with hot-reload enabled for both frontend and backend.

### Building

**Development Build**
```bash
wails build
```

**Production Build with NSIS Installer (Windows)**
```bash
wails build -nsis
```

The built application will be in `build/bin/`.


## Database Location

Clips are stored in a SQLite database at:
```
Windows: %APPDATA%\clipussy\db\gyatt.db
```


## Keyboard Shortcuts

| Shortcut | Action |
|---|---|
| `Ctrl + Shift + V` | Summon / show Clipcat from any application (system-wide) |
| `Ctrl + F` | Focus the search bar |
| `Alt + M` | Toggle Mini Clip mode |
| `Alt + H` | Toggle Privacy Mode (hide content) |
| `Alt + S` | Toggle sound effects |


## Customization

### Changing Clip Limit
The storage limit is dynamic and stored in the database. Adjust it in the settings panel (100-500, steps of 50) or programmatically:

```typescript
import { GetStorageLimit, UpdateStorageLimit } from './wailsjs/go/main/App'

const limit = await GetStorageLimit()
await UpdateStorageLimit(200)
```

Or directly in the database:
```sql
INSERT OR REPLACE INTO clip_storage_limit (id, limit_count) VALUES (0, 200);
```

### Blocking an Application
Add its process name in the "Blocked Apps" settings panel, or directly:
```sql
INSERT OR IGNORE INTO ignore_list (process_name) VALUES ('1password.exe');
```

### Adjusting Sound Volume
Edit the volume argument in component handlers:
```typescript
playSound("/sounds/file.mp3", soundOn, 0.3)  // 0.0 to 1.0
```

### Modifying UI Colors
Edit `frontend/src/index.css` and Tailwind classes in components.


## Contributing

Contributions are welcome! Feel free to submit issues and pull requests.


## Author

**Onyekwelu Jesse** ([@d3uceY](https://github.com/d3uceY))


## License

This project is licensed under the MIT License.


## Acknowledgments

- [Wails](https://wails.io/) for the amazing Go + Web framework
- All open-source contributors whose libraries made this possible

---

Made with love by d3uceY
