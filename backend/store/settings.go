package store

func GetGhostMode() (bool, error) {
	var v int
	err := DB.QueryRow(`SELECT ghost_mode FROM settings WHERE id = 0`).Scan(&v)
	if err != nil {
		return false, err
	}
	return v == 1, nil
}

func SetGhostMode(enabled bool) error {
	val := 0
	if enabled {
		val = 1
	}
	_, err := DB.Exec(`UPDATE settings SET ghost_mode = ? WHERE id = 0`, val)
	return err
}
