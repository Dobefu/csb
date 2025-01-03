package database

import (
	"os"
	"testing"

	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	var err error

	err = Connect()
	assert.NotEqual(t, nil, err)

	init_env.Main("../../.env.test")

	err = Connect()
	assert.Equal(t, nil, err)

	oldDbType := os.Getenv("DB_TYPE")
	os.Setenv("DB_TYPE", "bogus")

	err = Connect()
	assert.NotEqual(t, nil, err)

	os.Setenv("DB_TYPE", oldDbType)
}

func TestGetConnectionDetails(t *testing.T) {
	var err error

	init_env.Main("../../.env.test")

	oldDbType := os.Getenv("DB_TYPE")
	os.Setenv("DB_TYPE", "")

	_, _, err = getConnectionDetails()
	assert.NotEqual(t, nil, err)

	os.Setenv("DB_TYPE", oldDbType)
	_, _, err = getConnectionDetails()
	assert.Equal(t, nil, err)
}
