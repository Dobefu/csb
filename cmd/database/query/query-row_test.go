package query

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Dobefu/csb/cmd/database/structs"
	"github.com/stretchr/testify/assert"
)

func TestQueryRowMysql(t *testing.T) {
	mock, cleanup := setupTest(t, "mysql")
	defer cleanup()

	fields := []string{"id", "path"}
	where := []structs.QueryWhere{
		{Name: "id", Operator: structs.EQUALS, Value: 1},
	}

	expectedSQL := "SELECT id, path FROM routes WHERE id = \\? LIMIT 1"
	rows := sqlmock.NewRows([]string{"id", "path"}).AddRow(1, "/home")
	(*mock).ExpectQuery(expectedSQL).
		WithArgs(1).
		WillReturnRows(rows)

	result := QueryRow("routes", fields, where)
	assert.NotNil(t, result)

	var id int
	var path string
	err := result.Scan(&id, &path)
	assert.NoError(t, err)
	assert.Equal(t, 1, id)
	assert.Equal(t, "/home", path)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestQueryRowSqlite3(t *testing.T) {
	mock, cleanup := setupTest(t, "sqlite3")
	defer cleanup()

	fields := []string{"id", "key", "value"}
	where := []structs.QueryWhere{
		{Name: "key", Operator: structs.EQUALS, Value: "welcome"},
	}

	expectedSQL := "SELECT id, key, value FROM translations WHERE key = \\? LIMIT 1"
	rows := sqlmock.NewRows([]string{"id", "key", "value"}).AddRow(1, "welcome", "Welcome")
	(*mock).ExpectQuery(expectedSQL).
		WithArgs("welcome").
		WillReturnRows(rows)

	result := QueryRow("translations", fields, where)
	assert.NotNil(t, result)

	var id int
	var key, value string
	err := result.Scan(&id, &key, &value)
	assert.NoError(t, err)
	assert.Equal(t, 1, id)
	assert.Equal(t, "welcome", key)
	assert.Equal(t, "Welcome", value)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestQueryRowPostgres(t *testing.T) {
	mock, cleanup := setupTest(t, "postgres")
	defer cleanup()

	fields := []string{"id", "path"}
	where := []structs.QueryWhere{
		{Name: "id", Operator: structs.EQUALS, Value: 1},
	}

	expectedSQL := "SELECT id, path FROM routes WHERE id = \\$1 LIMIT 1"
	rows := sqlmock.NewRows([]string{"id", "path"}).AddRow(1, "/home")
	(*mock).ExpectQuery(expectedSQL).
		WithArgs(1).
		WillReturnRows(rows)

	result := QueryRow("routes", fields, where)
	assert.NotNil(t, result)

	var id int
	var path string
	err := result.Scan(&id, &path)
	assert.NoError(t, err)
	assert.Equal(t, 1, id)
	assert.Equal(t, "/home", path)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestQueryRowNoWhere(t *testing.T) {
	mock, cleanup := setupTest(t, "mysql")
	defer cleanup()

	fields := []string{"id", "path"}
	var where []structs.QueryWhere

	expectedSQL := "SELECT id, path FROM routes LIMIT 1"
	rows := sqlmock.NewRows([]string{"id", "path"}).AddRow(1, "/home")
	(*mock).ExpectQuery(expectedSQL).
		WillReturnRows(rows)

	result := QueryRow("routes", fields, where)
	assert.NotNil(t, result)

	var id int
	var path string
	err := result.Scan(&id, &path)
	assert.NoError(t, err)
	assert.Equal(t, 1, id)
	assert.Equal(t, "/home", path)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestQueryRowUnsupportedDB(t *testing.T) {
	_, cleanup := setupTest(t, "bogus")
	defer cleanup()

	fields := []string{"id", "path"}
	where := []structs.QueryWhere{
		{Name: "id", Operator: structs.EQUALS, Value: 1},
	}

	result := QueryRow("routes", fields, where)
	assert.Nil(t, result)
}
