package database

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() error {
	connString := os.Getenv("DB_CONN")

	if connString == "" {
		return errors.New("DB_CONN is not set")
	}

	dbType := os.Getenv("DB_TYPE")

	if dbType == "" {
		return errors.New("DB_TYPR is not set")
	}

	db, err := sql.Open(dbType, connString)

	if err != nil {
		return err
	}

	DB = db
	return nil
}
