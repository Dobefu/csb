package api

import (
	"os"
	"testing"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/Dobefu/csb/cmd/migrate_db"
	"github.com/stretchr/testify/assert"
)

func TestGetEntryByUid(t *testing.T) {
	init_env.Main("../../.env.test")
	err := database.Connect()
	assert.Equal(t, nil, err)

	err = migrate_db.Main(true)
	assert.Equal(t, nil, err)

	oldDb := os.Getenv("DB_CONN")
	os.Setenv("DB_CONN", "file:/")
	err = database.Connect()
	assert.Equal(t, nil, err)

	_, err = GetEntryByUid("", "", false)
	assert.NotEqual(t, nil, err)

	os.Setenv("DB_CONN", oldDb)
	err = database.Connect()
	assert.Equal(t, nil, err)

	err = insertPage("testingen", "testing", "")
	assert.Equal(t, nil, err)

	_, err = GetEntryByUid("testing", "en", true)
	assert.Equal(t, nil, err)
}
