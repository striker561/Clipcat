<p align="center">
  <img src="./build/appicon.png" alt="icon" width="90">
</p>
<div align="center">
<h1>Clipcat</h1>
</div>

A creative and stylish clipboard manager built with Wails, designed to keep track of everything you copy through a clean, paper-aesthetic interface. It automatically records every clipboard change in real time, storing your copy history so you can easily revisit, reuse, and manage past content whenever you need it.

<img width="1912" height="1026" alt="image" src="https://github.com/user-attachments/assets/ec4870c1-8555-49e0-9cb4-52c83f7551b0" />


## ⬇️ Download
![Clipcat Banner](https://img.shields.io/badge/Made%20with-Wails-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/license-MIT-blue?style=for-the-badge)

[![Download Clipcat v0.7.1 for Windows](https://img.shields.io/badge/Download-Windows%20Installer-0078D4?style=for-the-badge&logo=windows&logoColor=white)](https://github.com/d3uceY/Clipcat/releases/download/v0.7.1/Clipcat.exe)

**[⬇️ Clipcat](https://github.com/d3uceY/Clipcat/releases/download/v0.7.1/Clipcat.exe)**

> Windows 10/11 (64-bit) | Version 0.7.1

![clipcat (1)](https://github.com/user-attachments/assets/ca28ae42-2a9d-42c3-9a47-183808d59cf6)

### ⚠️ Windows SmartScreen Warning

When running the app for the first time, Windows SmartScreen may show a warning because the app is not yet code-signed. This is normal for open-source applications.

**To run the app:**
1. Click "More info" on the SmartScreen warning
2. Click "Run anyway"

Or alternatively:
1. Right-click the downloaded .exe file
2. Select "Properties"
3. Check "Unblock" at the bottom
4. Click "Apply" → "OK"
5. Run the executable

The app is safe and [open source](https://github.com/d3uceY/Clipcat) - you can verify the code yourself!



## ✨ Features

- **Automatic Clipboard Monitoring** - Automatically captures everything you copy (text and images)
- **Image Support** - Captures and displays clipboard images with base64 encoding
- **Pin Important Clips** - Keep your most-used clips at the top
- <img width="577" height="471" alt="image" src="https://github.com/user-attachments/assets/aa6e7ff1-159e-4fc0-8565-fdc9208a07c4" />

- **Fast Search** - Quickly find clips with Ctrl+F
- <img width="530" height="112" alt="showcase-search" src="https://github.com/user-attachments/assets/c2c76d50-5f94-481d-9191-ad37f2518967" />

- **Unique Paper Aesthetic** - Beautiful hand-drawn, notebook-style UI
- **Easy Management** - Copy, pin, and delete clips with intuitive controls
- **Edit Clips** - Modify the content of your saved clips anytime
- **Manual Creation** - Add new clips directly from the app without copying
- <img width="705" height="324" alt="showcase-action-btns" src="https://github.com/user-attachments/assets/d03a6634-8b41-4d78-a976-662b4c2b8f89" />

- **Privacy Mode** - Instantly hide clip content for privacy or during screen sharing
- <img width="242" height="227" alt="image" src="https://github.com/user-attachments/assets/9fed26d4-31cb-4ff0-bdd4-e8ee25780dba" />

- <img width="1223" height="244" alt="image" src="https://github.com/user-attachments/assets/0d242347-2e1d-46bf-b57f-20255d7c4fd1" />

- **Bulk Actions** - Quickly delete all, recent, or pinned clips via the settings menu
- <img width="196" height="179" alt="image" src="https://github.com/user-attachments/assets/a6cb09b3-14fa-4bc7-884a-57fc0d17d561" />

- <img width="467" height="193" alt="image" src="https://github.com/user-attachments/assets/a78a8a32-0d91-4920-ba4a-4100bb8d8cca" />

- **Full Content View** - Click any clip to view complete content in a scrollable dialog
- <img width="572" height="552" alt="image" src="https://github.com/user-attachments/assets/d85b3231-0380-4643-952a-e8d5dfcc71c4" />

- **Duplicate Detection** - Automatically prevents saving duplicate clipboard content
- **Sound Effects** - Audible feedback for actions
- **Persistent Storage** - SQLite database keeps your clips safe
- **Configurable Storage Limit** - Customize how many clips to keep (default: 100)
- <img width="259" height="335" alt="image" src="https://github.com/user-attachments/assets/ca3ed1ef-8a44-4aa2-a759-37d025e0682b" />

- **Mini Clip Mode** - A compact, always-on-top window for unobtrusive usage
- <img width="270" height="330" alt="image" src="https://github.com/user-attachments/assets/ddedff44-5007-4f57-b98e-006e28e24e71" />

- **Startup Support** - Option to launch automatically when your system starts
- <img width="293" height="335" alt="image" src="https://github.com/user-attachments/assets/2437742c-48c4-4e09-9726-19551d86eb54" />

- **Auto Update Check** - Automatically checks for new versions on GitHub

## 🛠️ Technologies Used

### Backend
- **[Go](https://golang.org/)** - Core application logic
- **[Wails v2](https://wails.io/)** - Desktop application framework
- **[SQLite](https://www.sqlite.org/)** (via modernc.org/sqlite) - Local database for clip storage
- **[golang.design/x/clipboard](https://github.com/golang-design/clipboard)** - Cross-platform clipboard access with image support
- **Windows API** (lxn/win) - Native Windows clipboard monitoring

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
## 🏗️ Architecture

Clipcat follows a clean architecture pattern with clear separation between frontend and backend:

```
┌─────────────────────────────────────────────────┐
│                   Frontend (React)               │
│  ┌────────────┐  ┌──────────┐  ┌──────────┐    │
│  │  UI Layer  │  │ Context  │  │ Components│    │
│  │  (TSX/CSS) │  │ Provider │  │  (Cards)  │    │
│  └────────────┘  └──────────┘  └──────────┘    │
└──────────────────────┬──────────────────────────┘
                       │ Wails Bridge (IPC)
┌──────────────────────┴──────────────────────────┐
│                Backend (Go)                      │
│  ┌────────────┐  ┌──────────┐  ┌──────────┐    │
│  │   App.go   │  │ clips.go │  │  db.go   │    │
│  │  (Bridge)  │  │ (Logic)  │  │(Storage) │    │
│  └────────────┘  └──────────┘  └──────────┘    │
└──────────────────────┬──────────────────────────┘
                       │
          ┌────────────┴────────────┐
          │                         │
    ┌─────▼─────┐          ┌────────▼────────┐
    │  SQLite   │          │ OS Clipboard    │
    │ Database  │          │   Listener      │
    └───────────┘          └─────────────────┘
```

### Data Flow

1. **Clipboard Monitoring**
   - Windows clipboard listener runs in the background
   - Detects clipboard changes via Windows API
   - Filters out duplicate or empty content

2. **Data Storage**
   - New clips are saved to SQLite database
   - Automatic cleanup keeps only 100 most recent clips (prioritizing pinned)
   - Each clip stores: content, type, timestamp, and pinned status

3. **Frontend Updates**
   - Backend emits events when clipboard changes
   - React context manages clip state
   - UI automatically re-renders with new data

4. **User Actions**
   - Copy: Uses browser clipboard API
   - Pin/Unpin: Toggles database flag, reorders UI
   - Delete: Removes from database, refreshes list
   - Search: Client-side filtering with instant results

## 📂 Project Structure

```
Clipcat/
├── app.go                      # Main application entry point
├── clips.go                    # Clip CRUD operations
├── db.go                       # Database initialization
├── main.go                     # Wails runtime setup
├── go.mod                      # Go dependencies
├── wails.json                  # Wails configuration
├── internal/
│   └── clipboard/
│       └── listener_window.go  # Windows clipboard listener
├── frontend/
│   ├── src/
│   │   ├── App.tsx            # Root component
│   │   ├── components/
│   │   │   ├── page.tsx       # Main page layout
│   │   │   └── ui/
│   │   │       ├── clip-card.tsx   # Individual clip card
│   │   │       └── dialog.tsx      # Modal dialog
│   │   ├── context/
│   │   │   └── ClipContext.tsx     # Global state management
│   │   ├── helpers/
│   │   │   ├── formatTime.ts       # Date formatting
│   │   │   └── playSound.ts        # Audio feedback
│   │   └── types/
│   │       └── clip.ts             # TypeScript interfaces
│   ├── wailsjs/                    # Auto-generated Wails bindings
│   ├── public/                     # Static assets
│   ├── package.json
│   └── vite.config.ts
└── build/                          # Build configuration
    └── windows/
        └── installer/              # NSIS installer config
```

## 🔧 How It Works

### Backend Implementation

**Clipboard Monitoring** (`internal/clipboard/listener_window.go`)
```go
// Polls clipboard every 500ms using Windows API
// Compares clipboard sequence numbers to detect changes
// Invokes callback when new content is detected
```

**Database Schema** (`db.go`)
```sql
CREATE TABLE clips (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    content TEXT,
    image BLOB,
    type TEXT NOT NULL,
    pinned BOOLEAN DEFAULT 0,
    created_at DATETIME
);

CREATE TABLE clip_storage_limit (
    id INTEGER PRIMARY KEY CHECK (id = 0),
    limit_count INTEGER DEFAULT 100
);
```

**Note:** The `content` field is nullable to support image-only clips. Either `content` or `image` will be populated based on the clip `type`.

**Key Operations** (`clips.go`)
- `getClips()` - Fetches all clips ordered by pinned status, then by date
- `clipExists()` - Checks if content already exists to prevent duplicates (text only)
- `addClip()` - Inserts new text clip (skips duplicates) and maintains dynamic clip limit
- `addImageClip()` - Inserts new image clip as BLOB and maintains dynamic clip limit
- `togglePinClip()` - Toggles pinned status by ID
- `deleteClip()` - Removes clip from database
- `getStorageLimit()` - Retrieves current storage limit from database
- `updateStorageLimit()` - Updates the maximum number of clips to store

### Frontend Implementation

**State Management** (`ClipContext.tsx`)
- Global state using React Context API
- Splits clips into pinned and recent arrays
- Listens for clipboard events from backend
- Provides `getClips()` method for manual refresh

**UI Components**
- **ClipCard** - Individual clip with copy/pin/delete actions; displays text or image based on type
- **Page** - Main layout with search, pinned section, recent section
- **AboutDialog** - Modal with app information and automatic update checking

**Animations** (GSAP)
- Paper curtain reveal on startup
- Cat character entrance
- Info button nudge animation
- Sound effects on interactions

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

## 📝 Database Location

Clips are stored in a SQLite database at:
```
Windows: %APPDATA%\clipussy\db\gyatt.db
```

## ⌨️ Keyboard Shortcuts

- `Ctrl + F` - Focus search bar
- `Ctrl + C` - Copy selected text (triggers clipboard monitoring)

## 🎨 Customization

### Changing Clip Limit
The storage limit is now dynamic and stored in the database. You can update it programmatically:

**From Frontend:**
```typescript
import { GetStorageLimit, UpdateStorageLimit } from './wailsjs/go/main/App'

// Get current limit
const limit = await GetStorageLimit()

// Set new limit
await UpdateStorageLimit(200)
```

**Or manually in the database:**
```sql
INSERT OR REPLACE INTO clip_storage_limit (id, limit_count) VALUES (0, 200);
```

### Adjusting Sound Volume
Edit respective handlers in `clip-card.tsx` and `page.tsx`:
```typescript
playSound("/sounds/file.mp3", soundOn, 0.3)  // 0.0 to 1.0
```

### Modifying UI Colors
Edit `frontend/src/index.css` and Tailwind classes in components.

##  Contributing

Contributions are welcome! Feel free to submit issues and pull requests.

##  Author

**Onyekwelu Jesse** ([@d3uceY](https://github.com/d3uceY))

##  License

This project is licensed under the MIT License.

##  Acknowledgments

- [Wails](https://wails.io/) for the amazing Go + Web framework
- All open-source contributors whose libraries made this possible

---

Made with 💜 by d3uceY
