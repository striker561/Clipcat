package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

// to install sqlite, use cli command <<< go get modernc.org/sqlite >>>

var DB *sql.DB

func InitDB(path string) error {
	var err error
	DB, err = sql.Open("sqlite", path)
	if err != nil {
		return err
	}

	fmt.Println("DB initialized on Bro")

	// ping this shit
	return DB.Ping()
}

func createTables() {
	query := `
	CREATE TABLE IF NOT EXISTS clips (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		image BLOB,
		type TEXT NOT NULL,
		pinned BOOLEAN DEFAULT 0,
		created_at DATETIME
	);

	CREATE TABLE IF NOT EXISTS clip_storage_limit (
		id INTEGER PRIMARY KEY CHECK (id = 0),
		limit_count INTEGER DEFAULT 100
	);

	CREATE TABLE IF NOT EXISTS ignore_list (
		process_name TEXT PRIMARY KEY
	);

	CREATE TABLE IF NOT EXISTS settings (
		id INTEGER PRIMARY KEY CHECK (id = 0),
		ghost_mode INTEGER DEFAULT 0
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		fmt.Printf("SQL Error: %v\nQuery: %s\n", err, query)
		panic(err)
	}
}

// this function adds the image column to the clips table if it doesn't exist
// if it already exists, it does nothing
func migrateClipsTable() {
	_, _ = DB.Exec(`ALTER TABLE clips ADD COLUMN image BLOB`)
}

// migrateSettingsTable seeds the single-row settings record for users whose
// DB was created before the settings table was introduced.
func migrateSettingsTable() {
	_, _ = DB.Exec(`INSERT OR IGNORE INTO settings (id, ghost_mode) VALUES (0, 0)`)
}
