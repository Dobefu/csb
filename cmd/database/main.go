package database

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() error {
	connString, dbType, err := getConnectionDetails()

	if err != nil {
		return err
	}

	db, err := sql.Open(dbType, connString)

	if err != nil {
		return err
	}

	DB = db
	return nil
}

func getConnectionDetails() (string, string, error) {
	connString := os.Getenv("DB_CONN")

	if connString == "" {
		return "", "", errors.New("DB_CONN is not set")
	}

	dbType := os.Getenv("DB_TYPE")

	if dbType == "" {
		return "", "", errors.New("DB_TYPE is not set")
	}

	return connString, dbType, nil
}
