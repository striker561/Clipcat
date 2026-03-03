package main

import (
	"fmt"
	"strings"
)

//
// Ignore List  process names excluded from clipboard capture
//
//
// Stored in the ignore_list table as lowercase exe names.
// e.g. "1password.exe", "keepassxc.exe"

func getIgnoreList() ([]string, error) {
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

func addIgnoreEntry(name string) error {
	name = strings.TrimSpace(strings.ToLower(name))
	if name == "" {
		return fmt.Errorf("addIgnoreEntry: process name cannot be empty")
	}
	_, err := DB.Exec("INSERT OR IGNORE INTO ignore_list (process_name) VALUES (?)", name)
	return err
}

func removeIgnoreEntry(name string) error {
	_, err := DB.Exec("DELETE FROM ignore_list WHERE process_name = ?", strings.ToLower(name))
	return err
}
