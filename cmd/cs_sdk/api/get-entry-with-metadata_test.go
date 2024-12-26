package api

import (
	"os"
	"testing"

	api_structs "github.com/Dobefu/csb/cmd/api/structs"
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/Dobefu/csb/cmd/migrate_db"
	"github.com/Dobefu/csb/cmd/remote_sync"
	"github.com/stretchr/testify/assert"
)

func TestGetEntryWithMetadata(t *testing.T) {
	var entry interface{}
	var altLocales []api_structs.AltLocale
	var breadcrumbs interface{}
	var err error

	var altLocalesEmpty []api_structs.AltLocale
	var breadcrumbsEmpty []structs.Route

	init_env.Main("../../../.env.test")
	err = database.Connect()
	assert.Equal(t, nil, err)

	err = migrate_db.Main(true)
	assert.Equal(t, nil, err)

	err = remote_sync.Sync(true)
	assert.Equal(t, nil, err)

	entry, altLocales, breadcrumbs, err = GetEntryWithMetadata(structs.Route{
		Uid:         "blt0617c28651fb44bf",
		ContentType: "basic_page",
		Locale:      "en",
	})
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, entry)
	assert.NotEqual(t, nil, altLocales)
	assert.NotEqual(t, nil, breadcrumbs)

	entry, altLocales, breadcrumbs, err = GetEntryWithMetadata(structs.Route{
		Uid:         "bogus",
		ContentType: "basic_page",
		Locale:      "en",
	})
	assert.NotEqual(t, nil, err)
	assert.Equal(t, nil, entry)
	assert.Equal(t, altLocalesEmpty, altLocales)
	assert.Equal(t, breadcrumbsEmpty, breadcrumbs)

	oldDb := os.Getenv("DB_CONN")
	os.Setenv("DB_CONN", "file:/")
	err = database.Connect()
	assert.Equal(t, nil, err)

	entry, altLocales, breadcrumbs, err = GetEntryWithMetadata(structs.Route{
		Uid:         "blt0617c28651fb44bf",
		ContentType: "basic_page",
		Locale:      "en",
	})
	assert.NotEqual(t, nil, err)
	assert.Equal(t, nil, entry)
	assert.Equal(t, altLocalesEmpty, altLocales)
	assert.Equal(t, breadcrumbsEmpty, breadcrumbs)

	os.Setenv("DB_CONN", oldDb)
	err = database.Connect()
	assert.Equal(t, nil, err)
}
