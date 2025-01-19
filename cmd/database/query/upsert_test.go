package query

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Dobefu/csb/cmd/database/structs"
	"github.com/stretchr/testify/assert"
)

func TestUpsertMysql(t *testing.T) {
	mock, cleanup := setupTest(t, "mysql")
	defer cleanup()

	routeValues := []structs.QueryValue{
		{Name: "id", Value: 1},
		{Name: "path", Value: "/"},
	}

	expectedRouteSQL := "INSERT INTO routes \\(id, path\\) VALUES \\(\\?, \\?\\) ON DUPLICATE KEY UPDATE id = VALUES\\(id\\), path = VALUES\\(path\\)"
	(*mock).ExpectExec(expectedRouteSQL).
		WithArgs(1, "/").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := Upsert("routes", routeValues)
	assert.NoError(t, err)

	translationValues := []structs.QueryValue{
		{Name: "id", Value: 1},
		{Name: "key", Value: "welcome"},
		{Name: "value", Value: "Welcome"},
	}

	expectedTranslationSQL := "INSERT INTO translations \\(id, key, value\\) VALUES \\(\\?, \\?, \\?\\) ON DUPLICATE KEY UPDATE id = VALUES\\(id\\), key = VALUES\\(key\\), value = VALUES\\(value\\)"
	(*mock).ExpectExec(expectedTranslationSQL).
		WithArgs(1, "welcome", "Welcome").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = Upsert("translations", translationValues)
	assert.NoError(t, err)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestUpsertMysqlError(t *testing.T) {
	mock, cleanup := setupTest(t, "mysql")
	defer cleanup()

	values := []structs.QueryValue{
		{Name: "id", Value: 1},
		{Name: "path", Value: "/"},
	}

	expectedSQL := "INSERT INTO routes \\(id, path\\) VALUES \\(\\?, \\?\\) ON DUPLICATE KEY UPDATE id = VALUES\\(id\\), path = VALUES\\(path\\)"
	(*mock).ExpectExec(expectedSQL).
		WithArgs(1, "/").
		WillReturnError(sqlmock.ErrCancelled)

	err := Upsert("routes", values)
	assert.Error(t, err)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestUpsertSqlite3(t *testing.T) {
	mock, cleanup := setupTest(t, "sqlite3")
	defer cleanup()

	values := []structs.QueryValue{
		{Name: "id", Value: 1},
		{Name: "key", Value: "welcome"},
		{Name: "value", Value: "Welcome"},
	}

	expectedSQL := "INSERT INTO translations \\(id, key, value\\) VALUES \\(\\?, \\?, \\?\\) ON CONFLICT DO UPDATE SET id = excluded.id, key = excluded.key, value = excluded.value"
	(*mock).ExpectExec(expectedSQL).
		WithArgs(1, "welcome", "Welcome").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := Upsert("translations", values)
	assert.NoError(t, err)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestUpsertSqlite3Error(t *testing.T) {
	mock, cleanup := setupTest(t, "sqlite3")
	defer cleanup()

	values := []structs.QueryValue{
		{Name: "id", Value: 1},
		{Name: "key", Value: "welcome"},
		{Name: "value", Value: "Welcome"},
	}

	expectedSQL := "INSERT INTO translations \\(id, key, value\\) VALUES \\(\\?, \\?, \\?\\) ON CONFLICT DO UPDATE SET id = excluded.id, key = excluded.key, value = excluded.value"
	(*mock).ExpectExec(expectedSQL).
		WithArgs(1, "welcome", "Welcome").
		WillReturnError(sqlmock.ErrCancelled)

	err := Upsert("translations", values)
	assert.Error(t, err)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestUpsertPostgres(t *testing.T) {
	mock, cleanup := setupTest(t, "postgres")
	defer cleanup()

	values := []structs.QueryValue{
		{Name: "id", Value: 1},
		{Name: "path", Value: "/"},
	}

	expectedSQL := "INSERT INTO routes \\(id, path\\) VALUES \\(\\$1, \\$2\\) ON CONFLICT ON CONSTRAINT routes_pkey DO UPDATE SET id = EXCLUDED.id, path = EXCLUDED.path"
	(*mock).ExpectExec(expectedSQL).
		WithArgs(1, "/").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := Upsert("routes", values)
	assert.NoError(t, err)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestUpsertPostgresError(t *testing.T) {
	mock, cleanup := setupTest(t, "postgres")
	defer cleanup()

	values := []structs.QueryValue{
		{Name: "id", Value: 1},
		{Name: "path", Value: "/"},
	}

	expectedSQL := "INSERT INTO routes \\(id, path\\) VALUES \\(\\$1, \\$2\\) ON CONFLICT ON CONSTRAINT routes_pkey DO UPDATE SET id = EXCLUDED.id, path = EXCLUDED.path"
	(*mock).ExpectExec(expectedSQL).
		WithArgs(1, "/").
		WillReturnError(sqlmock.ErrCancelled)

	err := Upsert("routes", values)
	assert.Error(t, err)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestUpsertUnsupportedDB(t *testing.T) {
	mock, cleanup := setupTest(t, "bogus")
	defer cleanup()

	values := []structs.QueryValue{
		{Name: "id", Value: 1},
		{Name: "path", Value: "/"},
	}

	(*mock).ExpectExec("").WillReturnError(sqlmock.ErrCancelled)

	err := Upsert("routes", values)
	assert.Error(t, err)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}
