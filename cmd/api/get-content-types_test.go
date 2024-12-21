package api

import (
	"os"
	"testing"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/stretchr/testify/assert"
)

func TestGetContentTypes(t *testing.T) {
	init_env.Main("../../.env.test")

	oldApiKey := os.Getenv("CS_API_KEY")
	os.Setenv("CS_API_KEY", "bogus")
	err := database.Connect()
	assert.Equal(t, nil, err)

	_, err = GetContentTypes()
	assert.NotEqual(t, nil, err)

	os.Setenv("CS_API_KEY", oldApiKey)
	err = database.Connect()
	assert.Equal(t, nil, err)

	contentTypes, err := GetContentTypes()
	assert.Equal(t, nil, err)

	assert.NotEqual(t, nil, contentTypes)
}
