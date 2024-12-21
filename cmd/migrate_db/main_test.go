package migrate_db

import (
	"os"
	"testing"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	var err error

	init_env.Main("../../.env.test")
	err = database.Connect()
	assert.Equal(t, nil, err)

	err = Main(false)
	assert.Equal(t, nil, err)

	err = Main(false)
	assert.Equal(t, nil, err)

	err = Main(true)
	assert.Equal(t, nil, err)

	err = down()
	assert.Equal(t, nil, err)

	oldDb := os.Getenv("DB_CONN")
	os.Setenv("DB_CONN", "file:/")
	err = database.Connect()
	assert.Equal(t, nil, err)

	err = Main(true)
	assert.NotEqual(t, nil, err)

	os.Setenv("DB_CONN", oldDb)
	err = database.Connect()
	assert.Equal(t, nil, err)
}

func TestUp(t *testing.T) {
	var err error

	err = up()
	assert.Equal(t, nil, err)

	err = up()
	assert.Equal(t, nil, err)

	err = down()
	assert.Equal(t, nil, err)

	oldDb := os.Getenv("DB_CONN")
	os.Setenv("DB_CONN", "file:/")
	err = database.Connect()
	assert.Equal(t, nil, err)

	err = up()
	assert.NotEqual(t, nil, err)

	os.Setenv("DB_CONN", oldDb)
	err = database.Connect()
	assert.Equal(t, nil, err)
}

func TestDown(t *testing.T) {
	var err error

	err = up()
	assert.Equal(t, nil, err)

	err = down()
	assert.Equal(t, nil, err)

	err = down()
	assert.Equal(t, nil, err)

	oldDb := os.Getenv("DB_CONN")
	os.Setenv("DB_CONN", "file:/")
	err = database.Connect()
	assert.Equal(t, nil, err)

	err = down()
	assert.NotEqual(t, nil, err)

	os.Setenv("DB_CONN", oldDb)
	err = database.Connect()
	assert.Equal(t, nil, err)
}

func TestRunMigration(t *testing.T) {
	var err error

	err = runMigration("bogus", 0)
	assert.NotEqual(t, nil, err)

	err = down()
	assert.Equal(t, nil, err)

	oldDb := os.Getenv("DB_CONN")
	os.Setenv("DB_CONN", "file:/")
	err = database.Connect()
	assert.Equal(t, nil, err)

	err = runMigration("bogus", 0)
	assert.NotEqual(t, nil, err)

	os.Setenv("DB_CONN", oldDb)
	err = database.Connect()
	assert.Equal(t, nil, err)
}
