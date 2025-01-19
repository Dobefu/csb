package query

import (
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/logger"
)

func setupTest(t *testing.T, dbType string) (*sqlmock.Sqlmock, func()) {
	logger.SetExitOnFatal(false)

	originalDBType := os.Getenv("DB_TYPE")
	os.Setenv("DB_TYPE", dbType)

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	database.DB = db

	return &mock, func() {
		logger.SetExitOnFatal(true)
		db.Close()
		os.Setenv("DB_TYPE", originalDBType)
	}
}
