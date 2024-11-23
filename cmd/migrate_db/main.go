package migrate_db

import (
	"github.com/Dobefu/csb/cmd/database"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Main(reset bool) error {
	driver, err := mysql.WithInstance(database.DB, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "mysql", driver)

	if err != nil {
		return err
	}

	if reset {
		m.Down()
	}

	err = m.Up()

	if err != nil {
		return err
	}

	return nil
}
