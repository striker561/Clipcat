package main

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Clip struct {
	ID        string  `json:"id"`
	Type      string  `json:"type"`
	Content   *string `json:"content,omitempty"`
	Image     *string `json:"image,omitempty"` // base64 PNG
	Length    int     `json:"length"`
	Pinned    bool    `json:"isPinned"`
	CreatedAt string  `json:"createdAt"`
}

func getStorageLimit() (int, error) {
	query := `SELECT limit_count FROM clip_storage_limit WHERE id = 0`
	var limit int
	err := DB.QueryRow(query).Scan(&limit)
	if err != nil {
		// If no limit is set, insert default of 100 and return it
		insertQuery := `INSERT OR IGNORE INTO clip_storage_limit (id, limit_count) VALUES (0, 100)`
		_, insertErr := DB.Exec(insertQuery)
		if insertErr != nil {
			return 100, fmt.Errorf("failed to initialize storage limit: %v", insertErr)
		}
		return 100, nil
	}
	return limit, nil
}

func updateStorageLimit(newLimit int) error {
	if newLimit < 1 {
		return fmt.Errorf("storage limit must be at least 1")
	}

	query := `INSERT OR REPLACE INTO clip_storage_limit (id, limit_count) VALUES (0, ?)`
	_, err := DB.Exec(query, newLimit)
	if err != nil {
		return fmt.Errorf("failed to update storage limit: %v", err)
	}
	return nil
}

func getClips() ([]Clip, error) {
	query := `
		SELECT id, content, image, type, pinned, created_at, encrypted
		FROM clips
		ORDER BY pinned DESC, created_at DESC
	`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var clips []Clip

	for rows.Next() {
		var (
			id        int
			content   sql.NullString
			image     []byte
			clipType  string
			pinned    bool
			createdAt string
			encrypted bool
		)

		err := rows.Scan(&id, &content, &image, &clipType, &pinned, &createdAt, &encrypted)
		if err != nil {
			return nil, err
		}

		clip := Clip{
			ID:        fmt.Sprintf("clip_%03d", id),
			Type:      clipType,
			Pinned:    pinned,
			CreatedAt: createdAt,
		}

		if clipType == "text" && content.Valid {
			if encrypted {
				plaintext, err := decryptText(content.String)
				if err == nil {
					clip.Content = &plaintext
					clip.Length = len(plaintext)
				} else {
					// Decryption failed — fall back to raw value so the row
					// is still visible rather than silently dropped.
					clip.Content = &content.String
					clip.Length = len(content.String)
				}
			} else {
				clip.Content = &content.String
				clip.Length = len(content.String)
			}
		}

		if clipType == "image" {
			imgBytes := image
			if encrypted {
				if dec, err := decryptData(image); err == nil {
					imgBytes = dec
				}
			}
			encoded := base64.StdEncoding.EncodeToString(imgBytes)
			clip.Image = &encoded
			clip.Length = len(imgBytes)
		}

		clips = append(clips, clip)
	}

	return clips, nil
}

func clipExists(content string) (bool, error) {
	hash := hashContent([]byte(content))
	// Match new encrypted clips by their HMAC hash, or legacy unencrypted rows
	// by their plaintext content (backward compatibility).
	query := `SELECT COUNT(*) FROM clips WHERE content_hash = ? OR (encrypted = 0 AND content = ?)`
	var count int
	err := DB.QueryRow(query, hash, content).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check if clip exists: %v", err)
	}
	return count > 0, nil
}

func imageClipExists(image []byte) (bool, error) {
	hash := hashContent(image)
	// Match new encrypted clips by their HMAC hash, or legacy unencrypted rows
	// by their raw image bytes (backward compatibility).
	query := `SELECT COUNT(*) FROM clips WHERE content_hash = ? OR (encrypted = 0 AND image = ?)`
	var count int
	err := DB.QueryRow(query, hash, image).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check if image clip exists: %v", err)
	}
	return count > 0, nil
}

func addClip(content string, clipType string) error {
	// Check if content already exists
	exists, err := clipExists(content)
	if err != nil {
		return fmt.Errorf("failed to check for duplicate: %v", err)
	}
	if exists {
		// Content already exists, skip insertion
		return nil
	}

	enc, err := encryptText(content)
	if err != nil {
		return fmt.Errorf("failed to encrypt clip: %v", err)
	}
	hash := hashContent([]byte(content))
	query := `INSERT INTO clips (content, content_hash, type, encrypted, created_at) VALUES (?, ?, ?, 1, datetime('now'))`
	_, err = DB.Exec(query, enc, hash, clipType)
	if err != nil {
		return fmt.Errorf("failed to insert clip: %v", err)
	}

	// Get the storage limit from database
	limit, err := getStorageLimit()
	if err != nil {
		return fmt.Errorf("failed to get storage limit: %v", err)
	}

	// Delete old clips, keeping only the most recent up to the limit (prioritizing pinned)
	deleteQuery := `
		DELETE FROM clips
		WHERE id NOT IN (
			SELECT id FROM clips
			ORDER BY pinned DESC, created_at DESC
			LIMIT ?
		)
	`
	_, err = DB.Exec(deleteQuery, limit)
	if err != nil {
		return fmt.Errorf("failed to delete old clips: %v", err)
	}

	return nil
}

func addManualClip(content string, pinned bool) error {
	// Check if content already exists
	exists, err := clipExists(content)
	if err != nil {
		return fmt.Errorf("failed to check for duplicate: %v", err)
	}
	if exists {
		// Content already exists, skip insertion
		return nil
	}

	enc, err := encryptText(content)
	if err != nil {
		return fmt.Errorf("failed to encrypt clip: %v", err)
	}
	hash := hashContent([]byte(content))
	query := `INSERT INTO clips (content, content_hash, type, pinned, encrypted, created_at) VALUES (?, ?, ?, ?, 1, datetime('now'))`
	_, err = DB.Exec(query, enc, hash, "text", pinned)
	if err != nil {
		return fmt.Errorf("failed to insert clip: %v", err)
	}

	// Get the storage limit from database
	limit, err := getStorageLimit()
	if err != nil {
		return fmt.Errorf("failed to get storage limit: %v", err)
	}

	// Delete old clips, keeping only the most recent up to the limit (prioritizing pinned)
	deleteQuery := `
		DELETE FROM clips
		WHERE id NOT IN (
			SELECT id FROM clips
			ORDER BY pinned DESC, created_at DESC
			LIMIT ?
		)
	`
	_, err = DB.Exec(deleteQuery, limit)
	if err != nil {
		return fmt.Errorf("failed to delete old clips: %v", err)
	}

	return nil
}

func addImageClip(img []byte) error {

	// Check if image already exists
	exists, err := imageClipExists(img)
	if err != nil {
		return fmt.Errorf("failed to check for duplicate image: %v", err)
	}
	if exists {
		// Image already exists, skip insertion
		return nil
	}

	enc, err := encryptData(img)
	if err != nil {
		return fmt.Errorf("failed to encrypt image clip: %v", err)
	}
	hash := hashContent(img)
	query := `INSERT INTO clips (image, content_hash, type, encrypted, created_at) VALUES (?, ?, ?, 1, datetime('now'))`
	_, err = DB.Exec(query, enc, hash, "image")
	if err != nil {
		return fmt.Errorf("failed to insert image clip: %v", err)
	}

	// Get the storage limit from database
	limit, err := getStorageLimit()
	if err != nil {
		return fmt.Errorf("failed to get storage limit: %v", err)
	}

	// Delete old clips, keeping only the most recent up to the limit (prioritizing pinned)
	deleteQuery := `
		DELETE FROM clips
		WHERE id NOT IN (
			SELECT id FROM clips
			ORDER BY pinned DESC, created_at DESC
			LIMIT ?
		)
	`
	_, err = DB.Exec(deleteQuery, limit)
	if err != nil {
		return fmt.Errorf("failed to delete old clips: %v", err)
	}
	return nil
}

func updateClipContent(clipID int, newContent string) error {
	enc, err := encryptText(newContent)
	if err != nil {
		return fmt.Errorf("failed to encrypt updated content: %v", err)
	}
	hash := hashContent([]byte(newContent))
	query := `UPDATE clips SET content = ?, content_hash = ?, encrypted = 1 WHERE id = ?`
	result, err := DB.Exec(query, enc, hash, clipID)
	if err != nil {
		return fmt.Errorf("failed to update clip content: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("clip with id %d not found", clipID)
	}

	return nil
}

func togglePinClip(clipID int) error {
	query := `UPDATE clips SET pinned = NOT pinned WHERE id = ?`
	result, err := DB.Exec(query, clipID)
	if err != nil {
		return fmt.Errorf("failed to toggle pin status: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("clip with id %d not found", clipID)
	}

	return nil
}

func deleteClip(clipID int) error {
	query := `DELETE FROM clips WHERE id = ?`
	result, err := DB.Exec(query, clipID)
	if err != nil {
		return fmt.Errorf("failed to delete clip: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("clip with id %d not found", clipID)
	}

	return nil
}

func deleteAllClips(context context.Context) error {
	res, err := runtime.MessageDialog(context, runtime.MessageDialogOptions{
		Type:          runtime.QuestionDialog,
		Title:         "Delete All Clips?",
		Message:       "Are you sure you want to delete all clips? This action cannot be undone.",
		DefaultButton: "Yes",
	},
	)

	if err != nil {
		return fmt.Errorf("failed to show confirmation dialog: %v", err)
	}

	if res != "Yes" {
		return nil
	}

	query := `DELETE FROM clips`
	_, err = DB.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to delete all clips: %v", err)
	}
	DB.Exec(`VACUUM`)
	return nil
}

func deletePinnedClips(context context.Context) error {
	res, err := runtime.MessageDialog(context, runtime.MessageDialogOptions{
		Type:          runtime.QuestionDialog,
		Title:         "Delete Pinned Clips?",
		Message:       "Are you sure you want to delete all pinned clips? This action cannot be undone.",
		DefaultButton: "Yes",
	},
	)

	if err != nil {
		return fmt.Errorf("failed to show confirmation dialog: %v", err)
	}

	if res != "Yes" {
		return nil
	}

	query := `DELETE FROM clips WHERE pinned = 1`
	_, err = DB.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to delete pinned clips: %v", err)
	}
	DB.Exec(`VACUUM`)
	return nil
}

func deleteUnpinnedClips(context context.Context) error {
	res, err := runtime.MessageDialog(context, runtime.MessageDialogOptions{
		Type:          runtime.QuestionDialog,
		Title:         "Delete Unpinned Clips?",
		Message:       "Are you sure you want to delete all unpinned clips? This action cannot be undone.",
		DefaultButton: "Yes",
	},
	)

	if err != nil {
		return fmt.Errorf("failed to show confirmation dialog: %v", err)
	}

	if res != "Yes" {
		return nil
	}

	query := `DELETE FROM clips WHERE pinned = 0`
	_, err = DB.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to delete unpinned clips: %v", err)
	}
	DB.Exec(`VACUUM`)
	return nil
}
