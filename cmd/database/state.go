package database

func GetState(name string) (string, error) {
	row := DB.QueryRow("SELECT value FROM state WHERE name = ?", name)

	var value string
	err := row.Scan(&value)

	if err != nil {
		return "", err
	}

	return value, nil
}

func SetState(name string, value string) error {
	_, err := DB.Exec(
		"INSERT INTO state (name, value) VALUES (?, ?) ON DUPLICATE KEY UPDATE value = VALUES(value)",
		name,
		value,
	)

	if err != nil {
		return err
	}

	return nil
}
