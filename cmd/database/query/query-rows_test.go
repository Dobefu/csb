package query

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Dobefu/csb/cmd/database/structs"
	"github.com/stretchr/testify/assert"
)

func TestQueryRowsMysql(t *testing.T) {
	mock, cleanup := setupTest(t, "mysql")
	defer cleanup()

	fields := []string{"id", "path"}
	where := []structs.QueryWhere{
		{Name: "id", Operator: structs.EQUALS, Value: 1},
	}

	expectedSQL := "SELECT id, path FROM routes WHERE id = \\?"
	rows := sqlmock.NewRows([]string{"id", "path"}).
		AddRow(1, "/home")
	(*mock).ExpectQuery(expectedSQL).
		WithArgs(1).
		WillReturnRows(rows)

	result, err := QueryRows("routes", fields, where)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	var routes []struct {
		ID   int
		Path string
	}

	for result.Next() {
		var route struct {
			ID   int
			Path string
		}
		err := result.Scan(&route.ID, &route.Path)
		assert.NoError(t, err)
		routes = append(routes, route)
	}

	assert.Len(t, routes, 1)
	assert.Equal(t, 1, routes[0].ID)
	assert.Equal(t, "/home", routes[0].Path)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestQueryRowsSqlite3(t *testing.T) {
	mock, cleanup := setupTest(t, "sqlite3")
	defer cleanup()

	fields := []string{"id", "key", "value"}
	where := []structs.QueryWhere{
		{Name: "id", Operator: structs.EQUALS, Value: 2},
	}

	expectedSQL := "SELECT id, key, value FROM translations WHERE id = \\?"
	rows := sqlmock.NewRows([]string{"id", "key", "value"}).
		AddRow(2, "goodbye", "Goodbye")
	(*mock).ExpectQuery(expectedSQL).
		WithArgs(2).
		WillReturnRows(rows)

	result, err := QueryRows("translations", fields, where)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	var translations []struct {
		ID    int
		Key   string
		Value string
	}

	for result.Next() {
		var translation struct {
			ID    int
			Key   string
			Value string
		}
		err := result.Scan(&translation.ID, &translation.Key, &translation.Value)
		assert.NoError(t, err)
		translations = append(translations, translation)
	}

	assert.Len(t, translations, 1)
	assert.Equal(t, 2, translations[0].ID)
	assert.Equal(t, "goodbye", translations[0].Key)
	assert.Equal(t, "Goodbye", translations[0].Value)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestQueryRowsPostgres(t *testing.T) {
	mock, cleanup := setupTest(t, "postgres")
	defer cleanup()

	fields := []string{"id", "path"}
	where := []structs.QueryWhere{
		{Name: "id", Operator: structs.EQUALS, Value: 1},
	}

	expectedSQL := "SELECT id, path FROM routes WHERE id = \\$1"
	rows := sqlmock.NewRows([]string{"id", "path"}).
		AddRow(1, "/home")
	(*mock).ExpectQuery(expectedSQL).
		WithArgs(1).
		WillReturnRows(rows)

	result, err := QueryRows("routes", fields, where)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	var routes []struct {
		ID   int
		Path string
	}

	for result.Next() {
		var route struct {
			ID   int
			Path string
		}
		err := result.Scan(&route.ID, &route.Path)
		assert.NoError(t, err)
		routes = append(routes, route)
	}

	assert.Len(t, routes, 1)
	assert.Equal(t, 1, routes[0].ID)
	assert.Equal(t, "/home", routes[0].Path)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestQueryRowsNoWhere(t *testing.T) {
	mock, cleanup := setupTest(t, "mysql")
	defer cleanup()

	fields := []string{"id", "path"}
	var where []structs.QueryWhere

	expectedSQL := "SELECT id, path FROM routes"
	rows := sqlmock.NewRows([]string{"id", "path"}).
		AddRow(1, "/home").
		AddRow(2, "/about")
	(*mock).ExpectQuery(expectedSQL).
		WillReturnRows(rows)

	result, err := QueryRows("routes", fields, where)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	var routes []struct {
		ID   int
		Path string
	}

	for result.Next() {
		var route struct {
			ID   int
			Path string
		}
		err := result.Scan(&route.ID, &route.Path)
		assert.NoError(t, err)
		routes = append(routes, route)
	}

	assert.Len(t, routes, 2)
	assert.Equal(t, 1, routes[0].ID)
	assert.Equal(t, "/home", routes[0].Path)
	assert.Equal(t, 2, routes[1].ID)
	assert.Equal(t, "/about", routes[1].Path)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestQueryRowsUnsupportedDB(t *testing.T) {
	_, cleanup := setupTest(t, "bogus")
	defer cleanup()

	fields := []string{"id", "path"}
	where := []structs.QueryWhere{
		{Name: "id", Operator: structs.EQUALS, Value: 1},
	}

	result, err := QueryRows("routes", fields, where)
	assert.NoError(t, err)
	assert.Nil(t, result)
}
