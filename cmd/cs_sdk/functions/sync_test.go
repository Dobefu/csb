package functions

import (
	"os"
	"testing"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/Dobefu/csb/cmd/migrate_db"
	"github.com/stretchr/testify/assert"
)

func TestSync(t *testing.T) {
	var err error

	init_env.Main("../../../.env.test")

	err = database.Connect()
	assert.Equal(t, nil, err)

	err = migrate_db.Main(true)
	assert.Equal(t, nil, err)

	err = Sync(true)
	assert.Equal(t, nil, err)

	err = Sync(false)
	assert.Equal(t, nil, err)

	oldDb := os.Getenv("DB_CONN")
	os.Setenv("DB_CONN", "file:/")
	err = database.Connect()
	assert.Equal(t, nil, err)

	err = Sync(true)
	assert.NotEqual(t, nil, err)

	err = Sync(false)
	assert.NotEqual(t, nil, err)

	os.Setenv("DB_CONN", oldDb)
	err = database.Connect()
	assert.Equal(t, nil, err)

	oldApiKey := os.Getenv("CS_API_KEY")
	os.Setenv("CS_API_KEY", "bogus")

	err = Sync(false)
	assert.NotEqual(t, nil, err)

	os.Setenv("CS_API_KEY", oldApiKey)
}
