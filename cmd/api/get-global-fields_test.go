package api

import (
	"os"
	"testing"

	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/stretchr/testify/assert"
)

func TestGetGlobalFields(t *testing.T) {
	var err error

	init_env.Main("../../.env.test")

	oldApiKey := os.Getenv("CS_API_KEY")
	os.Setenv("CS_API_KEY", "bogus")

	_, err = GetGlobalFields()
	assert.NotEqual(t, nil, err)

	os.Setenv("CS_API_KEY", oldApiKey)

	globalFields, err := GetGlobalFields()
	assert.Equal(t, nil, err)

	assert.NotEqual(t, nil, globalFields)
}