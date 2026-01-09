package main

import (
	"fmt"
)

type Clip struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	Length    int    `json:"length"`
	Pinned    bool   `json:"isPinned"`
	CreatedAt string `json:"createdAt"`
}

// get the storage limit from the database
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

// update the storage limit in the database
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

// this gets all clips
func getClips() ([]Clip, error) {
	query := `SELECT id, content, type, pinned, created_at FROM clips ORDER BY pinned DESC, created_at DESC`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query clips: %v", err)
	}
	defer rows.Close()

	var clips []Clip
	for rows.Next() {
		var id int
		var content, clipType, createdAt string
		var pinned bool
		err := rows.Scan(&id, &content, &clipType, &pinned, &createdAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan clip: %v", err)
		}

		clips = append(clips, Clip{
			ID:        fmt.Sprintf("clip_%03d", id),
			Content:   content,
			Length:    len(content),
			Pinned:    pinned,
			CreatedAt: createdAt,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating clips: %v", err)
	}

	return clips, nil
}

// check if a clip with the same content already exists in the database
func clipExists(content string) (bool, error) {
	query := `SELECT COUNT(*) FROM clips WHERE content = ?`
	var count int
	err := DB.QueryRow(query, content).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check if clip exists: %v", err)
	}
	return count > 0, nil
}

// this adds a new clip typeshit
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

	query := `INSERT INTO clips (content, type, created_at) VALUES (?, ?, datetime('now'))`
	_, err = DB.Exec(query, content, clipType)
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

// this pins/unpins a clip by toggling its pinned status
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

// this deletes a clip by its ID
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
