package api

import (
	"os"
	"testing"

	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/stretchr/testify/assert"
)

func TestGetLocales(t *testing.T) {
	init_env.Main("../../.env.test")

	var locales map[string]interface{}
	var err error

	locales, err = GetLocales()
	assert.Equal(t, nil, err)
	assert.NotEqual(t, 0, len(locales))

	oldApiKey := os.Getenv("CS_API_KEY")
	os.Setenv("CS_API_KEY", "bogus")

	locales, err = GetLocales()
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 0, len(locales))

	os.Setenv("CS_API_KEY", oldApiKey)
}
