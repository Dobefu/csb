package migrate_db

import (
	"embed"
	"fmt"
	"strings"

	"github.com/Dobefu/csb/cmd/database/query"
	"github.com/Dobefu/csb/cmd/database/structs"
	"github.com/Dobefu/csb/cmd/logger"
)

//go:embed migrations/*
var content embed.FS

func Main(reset bool) error {
	var err error

	if reset {
		logger.Info("Reverting existing migrations")
		err = down()

		if err != nil {
			return err
		}
	}

	logger.Info("Performing migrations")
	err = up()

	if err != nil {
		return err
	}

	return nil
}

func down() error {
	version, _, err := getMigrationState()

	if err != nil {
		return err
	}

	if version == 0 {
		logger.Info("Nothing to revert")
		return nil
	}

	files, err := content.ReadDir("migrations")

	if err != nil {
		return err
	}

	migrationIndex := version + 1

	for i := len(files) - 1; i >= 0; i-- {
		file := files[i]
		name := file.Name()

		if strings.Split(name, ".")[1] != "down" {
			continue
		}

		migrationIndex = i

		if migrationIndex > version {
			continue
		}

		logger.Info("Running migration: %s", name)
		err = runMigration(name, migrationIndex)

		if err != nil {
			return err
		}
	}

	err = setMigrationState(migrationIndex, false)

	if err != nil {
		return err
	}

	return nil
}

func up() error {
	version, _, err := getMigrationState()

	if err != nil {
		return err
	}

	files, err := content.ReadDir("migrations")

	if err != nil {
		return err
	}

	migrationIndex := 0

	for _, file := range files {
		name := file.Name()

		if strings.Split(name, ".")[1] != "up" {
			continue
		}

		migrationIndex += 1

		if migrationIndex <= version {
			continue
		}

		logger.Info("Running migration: %s", name)
		err = runMigration(name, migrationIndex)

		if err != nil {
			return err
		}
	}

	err = setMigrationState(migrationIndex, false)

	if err != nil {
		return err
	}

	return nil
}

func createMigrationsTable() error {
	_, err := query.QueryRaw(
		`CREATE TABLE IF NOT EXISTS migrations(
      version bigint NOT NULL PRIMARY KEY,
      dirty boolean NOT NULL
    );`,
	)

	if err != nil {
		return err
	}

	return nil
}

func getMigrationState() (int, bool, error) {
	err := createMigrationsTable()

	if err != nil {
		return 0, true, err
	}

	row := query.QueryRow("migrations", []string{"version", "dirty"}, nil)

	var version int
	var dirty bool

	err = row.Scan(&version, &dirty)

	// If nothing is found, the table is empty.
	// This is fine, since an initial migration might produce this result.
	// When this happens, default values should be returned.
	if err != nil {
		return 0, false, nil
	}

	return version, dirty, nil
}

func setMigrationState(version int, dirty bool) error {
	err := createMigrationsTable()

	if err != nil {
		return err
	}

	err = query.Truncate("migrations")

	if err != nil {
		return err
	}

	err = query.Insert("migrations", []structs.QueryValue{
		{
			Name:  "version",
			Value: version,
		},
		{
			Name:  "dirty",
			Value: dirty,
		},
	})

	if err != nil {
		return err
	}

	return nil
}

func runMigration(filename string, index int) error {
	queryBytes, err := content.ReadFile(fmt.Sprintf("migrations/%s", filename))

	if err != nil {
		_ = setMigrationState(index, true)
		return err
	}

	_, err = query.QueryRaw(string(queryBytes))

	if err != nil {
		_ = setMigrationState(index, true)
		return err
	}

	return nil
}
