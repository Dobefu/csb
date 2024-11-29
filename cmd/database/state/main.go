package state

import (
	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/database/query"
	"github.com/Dobefu/csb/cmd/database/structs"
)

func GetState(name string) (string, error) {
	row := query.QueryRow("state", []string{"value"}, []structs.QueryWhere{{Name: "name", Value: name}})

	var value string
	err := row.Scan(&value)

	if err != nil {
		return "", err
	}

	return value, nil
}

func SetState(name string, value string) error {
	_, err := database.DB.Exec(
		"REPLACE INTO state (name, value) VALUES (?, ?)",
		name,
		value,
	)

	if err != nil {
		return err
	}

	return nil
}
