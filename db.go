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

// migrateEncryptionColumns adds the columns and table required for at-rest
// encryption.  Safe to call repeatedly — ALTER TABLE and CREATE TABLE IF NOT
// EXISTS are idempotent.
func migrateEncryptionColumns() {
	_, _ = DB.Exec(`ALTER TABLE clips ADD COLUMN encrypted INTEGER DEFAULT 0`)
	_, _ = DB.Exec(`ALTER TABLE clips ADD COLUMN content_hash TEXT`)
	_, _ = DB.Exec(`
		CREATE TABLE IF NOT EXISTS encryption_meta (
			id          INTEGER PRIMARY KEY CHECK (id = 0),
			machine_key TEXT NOT NULL
		)
	`)
}

// migrateEncryptOldClips re-encrypts every pre-existing unencrypted row so
// that all clip data at rest is protected after the first run of a new version.
// Rows that fail to encrypt are left untouched (they will still be readable as
// plaintext via the backward-compatibility path in getClips).
func migrateEncryptOldClips() {
	type legacyRow struct {
		id       int
		content  sql.NullString
		image    []byte
		clipType string
	}

	rows, err := DB.Query(`SELECT id, content, image, type FROM clips WHERE encrypted = 0`)
	if err != nil {
		return
	}

	var clips []legacyRow
	for rows.Next() {
		var r legacyRow
		if err := rows.Scan(&r.id, &r.content, &r.image, &r.clipType); err == nil {
			clips = append(clips, r)
		}
	}
	rows.Close()

	for _, c := range clips {
		switch c.clipType {
		case "text":
			if !c.content.Valid || c.content.String == "" {
				continue
			}
			enc, err := encryptText(c.content.String)
			if err != nil {
				continue
			}
			hash := hashContent([]byte(c.content.String))
			_, _ = DB.Exec(
				`UPDATE clips SET content = ?, content_hash = ?, encrypted = 1 WHERE id = ?`,
				enc, hash, c.id,
			)
		case "image":
			if len(c.image) == 0 {
				continue
			}
			enc, err := encryptData(c.image)
			if err != nil {
				continue
			}
			hash := hashContent(c.image)
			_, _ = DB.Exec(
				`UPDATE clips SET image = ?, content_hash = ?, encrypted = 1 WHERE id = ?`,
				enc, hash, c.id,
			)
		}
	}
}
