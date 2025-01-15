package assets

import (
	"os"
	"testing"

	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/Dobefu/csb/cmd/migrate_db"
	"github.com/stretchr/testify/assert"
)

func TestSetAsset(t *testing.T) {
	init_env.Main("../../../.env.test")
	var err error

	oldDb := os.Getenv("DB_CONN")
	os.Setenv("DB_CONN", "file:/")
	err = database.Connect()
	assert.Equal(t, nil, err)

	err = SetAsset(structs.Asset{})
	assert.NotEqual(t, nil, err)

	os.Setenv("DB_CONN", oldDb)
	err = database.Connect()
	assert.Equal(t, nil, err)

	err = migrate_db.Main(true)
	assert.Equal(t, nil, err)

	err = SetAsset(structs.Asset{})
	assert.Equal(t, nil, err)
}
