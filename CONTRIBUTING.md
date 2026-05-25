# Contributing to Clipcat

Thanks for your interest in contributing! Below is everything you need to get the project running locally and submit changes.

---

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Project Structure](#project-structure)
3. [Architecture](#architecture)
4. [How It Works](#how-it-works)
5. [Setting Up Locally](#setting-up-locally)
6. [Development Workflow](#development-workflow)
7. [Building a Release Binary](#building-a-release-binary)
8. [Code Style](#code-style)
9. [Submitting a Pull Request](#submitting-a-pull-request)
10. [Reporting Bugs](#reporting-bugs)

---

## Prerequisites

| Tool | Version | Install |
|---|---|---|
| Go | ≥ 1.24 | https://go.dev/dl/ |
| Node.js | ≥ 18 | https://nodejs.org |
| npm | bundled with Node | — |
| Wails CLI | v2.12.0 | `go install github.com/wailsapp/wails/v2/cmd/wails@v2.12.0` |
| NSIS (installer only) | latest | https://nsis.sourceforge.io |

> Clipcat targets **Windows 10/11 (64-bit)** only. The clipboard listener, focus tracker, hotkey, and single-instance mutex all use Windows-only APIs. You must build and test on Windows.

---

## Project Structure

```
Clipcat/
├── app.go            # Wails App struct — startup, all frontend-exposed methods
├── tray.go           # Embed shim — holds the .ico asset, thin App receiver
├── main.go           # Wails runtime entry point
├── go.mod / go.sum   # Go module definition
├── wails.json        # Wails configuration (name, output filename, frontend scripts)
│
├── backend/
│   ├── store/        # package store — all SQLite logic
│   │   ├── db.go         DB init, table creation, schema migrations
│   │   ├── encrypt.go    AES-256-GCM encryption, key derivation, HMAC hashing
│   │   ├── clips.go      Clip type, CRUD operations, storage-limit enforcement
│   │   ├── settings.go   Ghost/Quick Paste mode flag
│   │   └── ignore.go     Blocked-app process list
│   ├── tray/         # package tray — systray logic
│   │   └── tray.go       tray.Start(icon, onShow, onQuit)
│   ├── platform/     # package platform — OS-level utilities
│   │   └── single_instance_windows.go   Mutex-based single-instance guard
│   └── lib/          # Shared low-level packages
│       ├── clipboard/    Windows clipboard listener + focus tracker + paste sim
│       └── startup/      Windows startup shortcut management
│
└── frontend/
    ├── src/
    │   ├── App.tsx
    │   ├── components/   UI components (ClipCard, page, dialogs, etc.)
    │   ├── context/      ClipContext — global React state
    │   ├── helpers/      Utility functions (formatTime, playSound, insertLinks…)
    │   ├── hooks/        Custom hooks (use-card-row-span, use-relative-time)
    │   └── types/        TypeScript interfaces (Clip)
    ├── wailsjs/          Auto-generated Wails Go→TS bindings (do not edit)
    └── public/           Static assets (sounds, cursors, textures)
```

---

## Architecture

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
|  +----------+  +--------------------+           |
|  |  app.go  |  |   backend/store/   |           |
|  | (Bridge) |  | clips  db  encrypt |           |
|  +----------+  | settings  ignore   |           |
|  +----------+  +--------------------+           |
|  |  tray.go |  +----------+  +--------------+   |
|  | (Shim)   |  |  backend/|  |   backend/   |   |
|  +----------+  |  tray/   |  |  platform/   |   |
|                +----------+  +--------------+   |
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

1. **Clipboard Monitoring** — A hidden `HWND_MESSAGE` window receives `WM_CLIPBOARDUPDATE` from Windows. A 150 ms debounce prevents duplicate saves. The ignore list filters blocked processes before anything is saved.

2. **Focus Tracking** — A background goroutine polls `GetForegroundWindow` every 150 ms, always keeping the last non-Clipcat window on hand so the paste button has a valid target.

3. **Data Storage** — Clips are saved to SQLite with HMAC-based duplicate detection. Automatic cleanup keeps only the most recent N clips, always preserving pinned ones.

4. **Frontend Updates** — The backend emits `clipboard:changed` events; React context re-fetches and re-renders.

5. **User Actions** — Copy (browser clipboard API), paste to window (focus prev window + simulate Ctrl+V), pin/delete (DB update + re-render), search (client-side filter).

---

## How It Works

### Clipboard Listener (`backend/lib/clipboard/listener_window.go`)

A hidden `HWND_MESSAGE` window is registered with `AddClipboardFormatListener`. Windows delivers `WM_CLIPBOARDUPDATE` when any app writes to the clipboard. The same window handles `WM_HOTKEY` for the global `Ctrl+Shift+V` shortcut via `RegisterHotKey`.

```go
case WM_CLIPBOARDUPDATE:
    // 150 ms debounce + pause check + ignore list check
    // then calls onChangeCallback()

case WM_HOTKEY:
    // Snapshot the current foreground window, then show Clipcat
    capturePreviousWindow()
    go onHotkeyCallback()
```

### Focus Tracker (`backend/lib/clipboard/window_utils.go`)

```go
func StartFocusTracker() {
    go func() {
        for {
            time.Sleep(150 * time.Millisecond)
            hwnd, _, _ := procGetForegroundWindow.Call()
            // skip our own PID, store prevHWND
        }
    }()
}
```

### Database Schema (`backend/store/db.go`)

```sql
CREATE TABLE clips (
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    content    TEXT,           -- nullable; populated for text clips
    image      BLOB,           -- populated for image clips
    type       TEXT NOT NULL,  -- 'text' or 'image'
    pinned     BOOLEAN DEFAULT 0,
    encrypted  INTEGER DEFAULT 0,
    content_hash TEXT,         -- HMAC-SHA256 for deduplication
    created_at DATETIME
);

CREATE TABLE clip_storage_limit (id INTEGER PRIMARY KEY CHECK (id = 0), limit_count INTEGER DEFAULT 100);
CREATE TABLE ignore_list        (process_name TEXT PRIMARY KEY);
CREATE TABLE settings           (id INTEGER PRIMARY KEY CHECK (id = 0), ghost_mode INTEGER DEFAULT 0);
CREATE TABLE encryption_meta    (id INTEGER PRIMARY KEY CHECK (id = 0), machine_key TEXT NOT NULL);
```

### Encryption (`backend/store/encrypt.go`)

All clip content is encrypted at rest with AES-256-GCM using a per-installation key stored in `encryption_meta`. Deduplication uses HMAC-SHA256 so duplicates can be detected without comparing plaintext.

### Store API (`backend/store/clips.go`)

| Function | Purpose |
|---|---|
| `GetClips()` | All clips, ordered pinned-first then newest |
| `AddClip()` / `AddImageClip()` | Insert + enforce storage limit |
| `AddManualClip()` | User-created clip with optional pin |
| `UpdateClipContent()` | Edit existing clip content |
| `TogglePinClip()` | Toggle pinned flag |
| `DeleteClip()` | Remove by ID |
| `DeleteAllClips()` / `DeletePinnedClips()` / `DeleteUnpinnedClips()` | Bulk delete with confirmation dialog |

### Frontend State (`ClipContext.tsx`)

- Splits clips into `pinned` and `recent` arrays
- Listens for `clipboard:changed` events from the backend
- Manages all settings: sound, privacy mode, mini clip, startup, pause, blocked apps, Quick Paste

---

## Setting Up Locally

```bash
# 1. Clone
git clone https://github.com/d3uceY/Clipcat.git
cd Clipcat

# 2. Install Go dependencies
go mod download

# 3. Install frontend dependencies
cd frontend
npm install
cd ..
```

---

## Development Workflow

```bash
wails dev
```

This launches the app with:
- Hot-reload on frontend changes (Vite)
- Automatic Go recompilation on backend changes
- A browser devtools window accessible at the URL printed in the terminal

**Useful dev flags:**

| Flag | Effect |
|---|---|
| `wails dev -loglevel debug` | Verbose Wails runtime logging |
| `wails dev -browser` | Also open a browser tab (layout debugging) |

### Performance Testing

`backend/store/clips.go` has a `SeedTestClips(n int)` function for inserting large batches of test clips. To use it, uncomment the call in `app.go`:

```go
// app.go — inside startup(), after MigrateEncryptOldClips()
// store.SeedTestClips(500) // PERF TEST: uncomment to insert 500 test clips
```

> **Always recomment this before committing.** It inserts into the live database on every startup.

---

## Building a Release Binary

**Portable `.exe`**
```bash
wails build -clean -o Clipcat-windows-amd64
# Output: build/bin/Clipcat-windows-amd64.exe
```

**NSIS installer** (requires NSIS installed and `makensis` on PATH)
```bash
wails build -clean -nsis -o Clipcat-windows-amd64
# Output: build/bin/Clipcat-windows-amd64-installer.exe
```

The CI release pipeline (`.github/workflows/release.yml`) runs the installer build automatically on any tag matching `v*.*.*`.

---

## Code Style

### Go

- Standard `gofmt` formatting — run `gofmt -w .` before committing
- Exported functions in `backend/` packages use `PascalCase`
- Internal (package-private) helpers stay `camelCase`
- Keep `app.go` as a thin coordinator — it should call `store.*` and `clipboard.*`, not contain business logic itself
- Do not add direct database calls to `app.go`; put them in `backend/store/`

### TypeScript / React

- Components use `.tsx`, pure utilities use `.ts`
- `React.memo` with explicit comparators on any component that renders inside a large list
- Avoid `useEffect` for derived state — prefer `useMemo`
- New hooks go in `frontend/src/hooks/`
- New UI-only components go in `frontend/src/components/`

---

## Submitting a Pull Request

1. Fork the repo and create a branch: `git checkout -b feat/your-feature`
2. Make your changes and verify the build: `go build ./...`
3. Run the app with `wails dev` and manually test the affected flows
4. Commit with a clear message: `feat: add X` / `fix: Y when Z` / `refactor: move X to Y`
5. Push and open a PR against `main`
6. Describe what changed and why in the PR description

### What gets reviewed

- Does `go build ./...` pass cleanly?
- Are new backend functions placed in the right package (`store`, `platform`, `lib/clipboard`, etc.)?
- Are there no direct SQL calls in `app.go`?
- Does the UI still look and feel correct on Windows?

---

## Reporting Bugs

Open an issue at https://github.com/d3uceY/Clipcat/issues with:

- Clipcat version (visible in the About dialog)
- Windows version (`winver`)
- Steps to reproduce
- Expected vs actual behaviour
- Any error output from the console (open DevTools with `F12` in dev mode, or check the Wails log)
