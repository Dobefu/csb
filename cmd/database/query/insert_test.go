package query

import (
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/database/structs"
	"github.com/Dobefu/csb/cmd/logger"
	"github.com/stretchr/testify/assert"
)

func setupInsertTest(t *testing.T, dbType string) (*sqlmock.Sqlmock, func()) {
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

func TestInsertMysql(t *testing.T) {
	mock, cleanup := setupInsertTest(t, "mysql")
	defer cleanup()

	routeValues := []structs.QueryValue{
		{Name: "id", Value: 1},
		{Name: "path", Value: "/"},
	}

	expectedRouteSQL := "INSERT INTO routes \\(id, path\\) VALUES \\(\\?, \\?\\)"
	(*mock).ExpectExec(expectedRouteSQL).
		WithArgs(1, "/").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := Insert("routes", routeValues)
	assert.NoError(t, err)

	translationValues := []structs.QueryValue{
		{Name: "id", Value: 1},
		{Name: "key", Value: "welcome"},
		{Name: "value", Value: "Welcome"},
	}

	expectedTranslationSQL := "INSERT INTO translations \\(id, key, value\\) VALUES \\(\\?, \\?, \\?\\)"
	(*mock).ExpectExec(expectedTranslationSQL).
		WithArgs(1, "welcome", "Welcome").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = Insert("translations", translationValues)
	assert.NoError(t, err)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestInsertMysqlError(t *testing.T) {
	mock, cleanup := setupInsertTest(t, "mysql")
	defer cleanup()

	values := []structs.QueryValue{
		{Name: "id", Value: 1},
		{Name: "path", Value: "/"},
	}

	expectedSQL := "INSERT INTO routes \\(id, path\\) VALUES \\(\\?, \\?\\)"
	(*mock).ExpectExec(expectedSQL).
		WithArgs(1, "/").
		WillReturnError(sqlmock.ErrCancelled)

	err := Insert("routes", values)
	assert.Error(t, err)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestInsertSqlite3(t *testing.T) {
	mock, cleanup := setupInsertTest(t, "sqlite3")
	defer cleanup()

	values := []structs.QueryValue{
		{Name: "id", Value: 1},
		{Name: "key", Value: "welcome"},
		{Name: "value", Value: "Welcome"},
	}

	expectedSQL := "INSERT INTO translations \\(id, key, value\\) VALUES \\(\\?, \\?, \\?\\)"
	(*mock).ExpectExec(expectedSQL).
		WithArgs(1, "welcome", "Welcome").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := Insert("translations", values)
	assert.NoError(t, err)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestInsertSqlite3Error(t *testing.T) {
	mock, cleanup := setupInsertTest(t, "sqlite3")
	defer cleanup()

	values := []structs.QueryValue{
		{Name: "id", Value: 1},
		{Name: "key", Value: "welcome"},
		{Name: "value", Value: "Welcome"},
	}

	expectedSQL := "INSERT INTO translations \\(id, key, value\\) VALUES \\(\\?, \\?, \\?\\)"
	(*mock).ExpectExec(expectedSQL).
		WithArgs(1, "welcome", "Welcome").
		WillReturnError(sqlmock.ErrCancelled)

	err := Insert("translations", values)
	assert.Error(t, err)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestInsertPostgres(t *testing.T) {
	mock, cleanup := setupInsertTest(t, "postgres")
	defer cleanup()

	values := []structs.QueryValue{
		{Name: "id", Value: 1},
		{Name: "path", Value: "/"},
	}

	expectedSQL := "INSERT INTO routes \\(id, path\\) VALUES \\(\\$1, \\$2\\)"
	(*mock).ExpectExec(expectedSQL).
		WithArgs(1, "/").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := Insert("routes", values)
	assert.NoError(t, err)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestInsertPostgresError(t *testing.T) {
	mock, cleanup := setupInsertTest(t, "postgres")
	defer cleanup()

	values := []structs.QueryValue{
		{Name: "id", Value: 1},
		{Name: "path", Value: "/"},
	}

	expectedSQL := "INSERT INTO routes \\(id, path\\) VALUES \\(\\$1, \\$2\\)"
	(*mock).ExpectExec(expectedSQL).
		WithArgs(1, "/").
		WillReturnError(sqlmock.ErrCancelled)

	err := Insert("routes", values)
	assert.Error(t, err)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestInsertUnsupportedDB(t *testing.T) {
	mock, cleanup := setupInsertTest(t, "bogus")
	defer cleanup()

	values := []structs.QueryValue{
		{Name: "id", Value: 1},
		{Name: "path", Value: "/"},
	}

	(*mock).ExpectExec("").WillReturnError(sqlmock.ErrCancelled)

	err := Insert("routes", values)
	assert.NoError(t, err)
}
