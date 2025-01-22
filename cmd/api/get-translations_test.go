package api

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Dobefu/csb/cmd/database/query"
	"github.com/Dobefu/csb/cmd/database/structs"
	"github.com/stretchr/testify/assert"
)

func setupGetTranslationsTest(t *testing.T) (sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	queryQueryRows = func(table string, fields []string, where []structs.QueryWhere) (*sql.Rows, error) {
		return db.Query("SELECT (.+) FROM translations", nil)
	}

	return mock, func() {
		db.Close()
		queryQueryRows = query.QueryRows
	}
}

func TestGetTranslationsSuccess(t *testing.T) {
	mock, cleanup := setupGetTranslationsTest(t)
	defer cleanup()

	cols := []string{"source", "translation", "category"}
	mock.ExpectQuery("SELECT (.+) FROM translations").WillReturnRows(
		sqlmock.NewRows(cols).AddRow("source", "translation", "category"),
	)

	translations, err := GetTranslations("en")
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"category.source": "translation"}, translations)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetTranslationsErrQuery(t *testing.T) {
	mock, cleanup := setupGetTranslationsTest(t)
	defer cleanup()

	mock.ExpectQuery("SELECT (.+) FROM translations").WillReturnError(
		sql.ErrNoRows,
	)

	translations, err := GetTranslations("en")
	assert.EqualError(t, err, sql.ErrNoRows.Error())
	assert.Equal(t, map[string]interface{}(nil), translations)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetTranslationsErrScan(t *testing.T) {
	mock, cleanup := setupGetTranslationsTest(t)
	defer cleanup()

	cols := []string{"source", "translation", "category"}
	mock.ExpectQuery("SELECT (.+) FROM translations").WillReturnRows(
		sqlmock.NewRows(cols).AddRow("source", "translation", nil),
	)

	translations, err := GetTranslations("en")
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{}, translations)
	assert.NoError(t, mock.ExpectationsWereMet())
}
