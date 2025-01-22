package state

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Dobefu/csb/cmd/database/query"
	"github.com/Dobefu/csb/cmd/database/structs"
	"github.com/stretchr/testify/assert"
)

func setupMockDB(t *testing.T) (sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	queryQueryRow = func(table string, fields []string, where []structs.QueryWhere) *sql.Row {
		return db.QueryRow("SELECT value FROM state WHERE name = ?", where[0].Value)
	}

	queryUpsert = func(table string, values []structs.QueryValue) error {
		_, err := db.Exec("INSERT INTO state", values[0].Value, values[1].Value)
		return err
	}

	cleanup := func() {
		db.Close()
		queryQueryRow = query.QueryRow
		queryUpsert = query.Upsert
	}

	return mock, cleanup
}

func TestGetStateSuccess(t *testing.T) {
	mock, cleanup := setupMockDB(t)
	defer cleanup()

	mock.ExpectQuery("SELECT value FROM state WHERE name = ?").
		WithArgs("test").
		WillReturnRows(sqlmock.NewRows([]string{"value"}).AddRow("test_value"))

	result, err := GetState("test")

	assert.NoError(t, err)
	assert.Equal(t, "test_value", result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetStateErr(t *testing.T) {
	mock, cleanup := setupMockDB(t)
	defer cleanup()

	mock.ExpectQuery("SELECT value FROM state WHERE name = ?").
		WithArgs("test").
		WillReturnError(sql.ErrNoRows)

	result, err := GetState("test")

	assert.Error(t, err)
	assert.Empty(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSetStateSuccess(t *testing.T) {
	mock, cleanup := setupMockDB(t)
	defer cleanup()

	mock.ExpectExec("INSERT INTO state").
		WithArgs("test", "test_value").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := SetState("test", "test_value")

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSetStateErr(t *testing.T) {
	mock, cleanup := setupMockDB(t)
	defer cleanup()

	mock.ExpectExec("INSERT INTO state").
		WithArgs("test", "test_value").
		WillReturnError(errors.New("database error"))

	err := SetState("test", "test_value")

	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
