package query

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Dobefu/csb/cmd/logger"
	"github.com/stretchr/testify/assert"
)

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
