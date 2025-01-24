package api

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/Dobefu/csb/cmd/database/query"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
	"github.com/stretchr/testify/assert"
)

func setupGetEntryByUrlTest(t *testing.T) (sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	queryQueryRow = func(table string, fields []string, where []db_structs.QueryWhere) *sql.Row {
		return db.QueryRow("SELECT (.+) FROM routes", nil)
	}

	return mock, func() {
		queryQueryRow = query.QueryRow
	}
}

func TestGetEntryByUrl(t *testing.T) {
	mock, cleanup := setupGetEntryByUrlTest(t)
	defer cleanup()

	cols := []string{"id", "uid", "title", "content_type", "locale", "slug", "url", "parent", "version", "updated_at", "exclude_sitemap", "published"}
	mock.ExpectQuery("SELECT (.+) FROM routes").WillReturnRows(
		sqlmock.NewRows(cols).AddRow("id", "uid", "title", "content_type", "locale", "slug", "url", "parent", 0, time.Time{}, true, true),
	)

	entry, err := GetEntryByUrl("/test-url", "test-locale", false)
	assert.NoError(t, err)
	assert.Equal(t, structs.Route{
		Id:             "id",
		Uid:            "uid",
		Title:          "title",
		ContentType:    "content_type",
		Locale:         "locale",
		Slug:           "slug",
		Url:            "url",
		Parent:         "parent",
		Version:        0,
		UpdatedAt:      time.Time{},
		ExcludeSitemap: true,
		Published:      true,
	}, entry)
}
