package api

import (
	"os"
	"testing"

	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/stretchr/testify/assert"
)

func TestGetContentType(t *testing.T) {
	var contentType map[string]interface{}
	var err error

	init_env.Main("../../.env.test")

	oldApiKey := os.Getenv("CS_API_KEY")
	os.Setenv("CS_API_KEY", "bogus")

	_, err = GetContentType("basic_page")
	assert.NotEqual(t, nil, err)

	os.Setenv("CS_API_KEY", oldApiKey)

	contentType, err = GetContentType("basic_page")
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, contentType)

	contentType, err = GetContentType("bogus")
	assert.NotEqual(t, nil, err)
	assert.NotEqual(t, nil, contentType)
}
