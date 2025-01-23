package api

import (
	"os"
	"testing"
	"time"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/database/query"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/Dobefu/csb/cmd/migrate_db"
	"github.com/stretchr/testify/assert"
)

func TestGetEntryByUrl(t *testing.T) {
	init_env.Main("../../.env.test")
	err := database.Connect()
	assert.Equal(t, nil, err)

	err = migrate_db.Main(true)
	assert.Equal(t, nil, err)

	oldDb := os.Getenv("DB_CONN")
	os.Setenv("DB_CONN", "file:/")
	err = database.Connect()
	assert.Equal(t, nil, err)

	_, err = GetEntryByUrl("", "", false)
	assert.NotEqual(t, nil, err)

	os.Setenv("DB_CONN", oldDb)
	err = database.Connect()
	assert.Equal(t, nil, err)

	err = insertPage("testingen", "testing", "")
	assert.Equal(t, nil, err)

	_, err = GetEntryByUrl("/testing", "en", true)
	assert.Equal(t, nil, err)
}

func insertPage(id string, uid string, parent string) error {
	return query.Insert("routes", []db_structs.QueryValue{
		{
			Name:  "id",
			Value: id,
		},
		{
			Name:  "uid",
			Value: uid,
		},
		{
			Name:  "title",
			Value: "Title",
		},
		{
			Name:  "content_type",
			Value: "basic_page",
		},
		{
			Name:  "locale",
			Value: "en",
		},
		{
			Name:  "slug",
			Value: "/testing",
		},
		{
			Name:  "url",
			Value: "/testing",
		},
		{
			Name:  "parent",
			Value: parent,
		},
		{
			Name:  "updated_at",
			Value: time.Now(),
		},
	})
}
