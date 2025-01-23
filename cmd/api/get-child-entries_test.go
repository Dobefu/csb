package api

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Dobefu/csb/cmd/database/query"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"

	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/stretchr/testify/assert"
)

func setupGetChildEntriesTest(t *testing.T) (sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	queryQueryRows = func(table string, fields []string, where []db_structs.QueryWhere) (*sql.Rows, error) {
		return db.Query("SELECT (.+) FROM routes", nil)
	}

	return mock, func() {
		db.Close()
		queryQueryRows = query.QueryRows
	}
}

func TestGetChildEntriesByUidSuccess(t *testing.T) {
	mock, cleanup := setupGetChildEntriesTest(t)
	defer cleanup()

	cols := []string{"id", "uid", "content_type", "locale", "slug", "url", "parent", "updated_at", "exclude_sitemap", "published"}
	mock.ExpectQuery("SELECT (.+) FROM routes").WillReturnRows(
		sqlmock.NewRows(cols).AddRow("id", "uid", "content_type", "locale", "slug", "url", "parent", time.Time{}, true, true),
	)

	routes, err := GetChildEntriesByUid("uid", "en", false)
	assert.NoError(t, err)
	assert.Equal(t, []structs.Route{{
		Id:             "id",
		Uid:            "uid",
		Title:          "",
		ContentType:    "content_type",
		Locale:         "locale",
		Slug:           "slug",
		Url:            "url",
		Parent:         "parent",
		Version:        0,
		UpdatedAt:      time.Time{},
		ExcludeSitemap: true,
		Published:      true,
	}}, routes)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetChildEntriesByUidErrQuery(t *testing.T) {
	mock, cleanup := setupGetChildEntriesTest(t)
	defer cleanup()

	mock.ExpectQuery("SELECT (.+) FROM routes").WillReturnError(sql.ErrNoRows)

	routes, err := GetChildEntriesByUid("uid", "en", false)
	assert.EqualError(t, err, sql.ErrNoRows.Error())
	assert.Equal(t, []structs.Route([]structs.Route{}), routes)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetChildEntriesByUidErrScan(t *testing.T) {
	mock, cleanup := setupGetChildEntriesTest(t)
	defer cleanup()

	cols := []string{"locacle"}
	mock.ExpectQuery("SELECT (.+) FROM routes").WillReturnRows(
		sqlmock.NewRows(cols).AddRow("locale"),
	)

	routes, err := GetChildEntriesByUid("uid", "en", false)
	assert.NoError(t, err)
	assert.Equal(t, []structs.Route([]structs.Route(nil)), routes)
	assert.NoError(t, mock.ExpectationsWereMet())
}
