package query

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Dobefu/csb/cmd/database"
	"github.com/stretchr/testify/assert"
)

func setupQueryRawTest(t *testing.T) (*sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	database.DB = db

	return &mock, func() {
		db.Close()
	}
}

func TestQueryRaw(t *testing.T) {
	mock, cleanup := setupQueryRawTest(t)
	defer cleanup()

	sql := "INSERT INTO users (routes) VALUES ('/')"
	(*mock).ExpectExec(sql).WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := QueryRaw(sql)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestQueryRawErr(t *testing.T) {
	mock, cleanup := setupQueryRawTest(t)
	defer cleanup()

	sql := "INSERT INTO non_existent_table (column) VALUES ('/')"
	(*mock).ExpectExec(sql).WillReturnError(sqlmock.ErrCancelled)

	result, err := QueryRaw(sql)
	assert.Error(t, err)
	assert.Nil(t, result)

	assert.NoError(t, (*mock).ExpectationsWereMet())
}
