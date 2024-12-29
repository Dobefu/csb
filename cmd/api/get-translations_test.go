package api

import (
	"os"
	"testing"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/database/query"
	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/Dobefu/csb/cmd/migrate_db"
	"github.com/stretchr/testify/assert"
)

func TestGetTranslations(t *testing.T) {
	var translations map[string]interface{}
	var err error

	init_env.Main("../../.env.test")

	err = database.Connect()
	assert.Equal(t, nil, err)

	err = migrate_db.Main(true)
	assert.Equal(t, nil, err)

	_, err = query.QueryRaw("INSERT INTO translations VALUES ('id', 'uid', 'source', 'translation', 'category', 'en')")
	assert.Equal(t, nil, err)

	translations, err = GetTranslations("en")
	assert.Equal(t, nil, err)
	assert.NotEqual(t, 0, len(translations))
	assert.Equal(t, "translation", translations["category.source"])

	oldDb := os.Getenv("DB_CONN")
	os.Setenv("DB_CONN", "file:/")
	err = database.Connect()
	assert.Equal(t, nil, err)

	translations, err = GetTranslations("en")
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 0, len(translations))

	os.Setenv("DB_CONN", oldDb)
	err = database.Connect()
	assert.Equal(t, nil, err)
}
