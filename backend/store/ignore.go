package store

import (
	"fmt"
	"strings"
)

func GetIgnoreList() ([]string, error) {
	rows, err := DB.Query("SELECT process_name FROM ignore_list ORDER BY process_name")
	if err != nil {
		return nil, fmt.Errorf("getIgnoreList: %w", err)
	}
	defer rows.Close()

	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	return names, nil
}

func AddIgnoreEntry(name string) error {
	name = strings.TrimSpace(strings.ToLower(name))
	if name == "" {
		return fmt.Errorf("addIgnoreEntry: process name cannot be empty")
	}
	_, err := DB.Exec("INSERT OR IGNORE INTO ignore_list (process_name) VALUES (?)", name)
	return err
}

func RemoveIgnoreEntry(name string) error {
	_, err := DB.Exec("DELETE FROM ignore_list WHERE process_name = ?", strings.ToLower(name))
	return err
}
