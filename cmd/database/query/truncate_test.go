package query

import (
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/logger"
	"github.com/stretchr/testify/assert"
)

func setupTest(t *testing.T, dbType string) (*sqlmock.Sqlmock, func()) {
	originalDBType := os.Getenv("DB_TYPE")
	os.Setenv("DB_TYPE", dbType)

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	database.DB = db

	return &mock, func() {
		db.Close()
		os.Setenv("DB_TYPE", originalDBType)
	}
}

func TestTruncateMysql(t *testing.T) {
	mock, cleanup := setupTest(t, "mysql")
	defer cleanup()

	(*mock).ExpectExec("TRUNCATE routes").WillReturnResult(sqlmock.NewResult(0, 0))

	err := Truncate("routes")
	assert.NoError(t, err)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestTruncateMysqlError(t *testing.T) {
	mock, cleanup := setupTest(t, "mysql")
	defer cleanup()

	(*mock).ExpectExec("TRUNCATE routes").WillReturnError(sqlmock.ErrCancelled)

	err := Truncate("routes")
	assert.Error(t, err)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestTruncateSqlite3(t *testing.T) {
	mock, cleanup := setupTest(t, "sqlite3")
	defer cleanup()

	(*mock).ExpectExec("DELETE FROM products").WillReturnResult(sqlmock.NewResult(0, 0))

	err := Truncate("products")
	assert.NoError(t, err)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestTruncateSqlite3Error(t *testing.T) {
	mock, cleanup := setupTest(t, "sqlite3")
	defer cleanup()

	(*mock).ExpectExec("DELETE FROM products").WillReturnError(sqlmock.ErrCancelled)

	err := Truncate("products")
	assert.Error(t, err)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestTruncatePostgres(t *testing.T) {
	mock, cleanup := setupTest(t, "postgres")
	defer cleanup()

	(*mock).ExpectExec("TRUNCATE translations").WillReturnResult(sqlmock.NewResult(0, 0))

	err := Truncate("translations")
	assert.NoError(t, err)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestTruncatePostgresError(t *testing.T) {
	mock, cleanup := setupTest(t, "postgres")
	defer cleanup()

	(*mock).ExpectExec("TRUNCATE translations").WillReturnError(sqlmock.ErrCancelled)

	err := Truncate("translations")
	assert.Error(t, err)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestTruncateUnsupportedDB(t *testing.T) {
	logger.SetExitOnFatal(false)
	_, cleanup := setupTest(t, "unsupported")
	defer cleanup()

	err := Truncate("test")
	assert.NoError(t, err)
}
