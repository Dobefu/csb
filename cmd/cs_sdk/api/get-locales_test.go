package api

import (
	"os"
	"testing"

	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/stretchr/testify/assert"
)

func TestGetLocales(t *testing.T) {
	init_env.Main("../../../.env.test")

	var locales interface{}

	locales = GetLocales()
	assert.NotEqual(t, nil, locales)

	oldApiKey := os.Getenv("CS_API_KEY")
	os.Setenv("CS_API_KEY", "bogus")

	locales = GetLocales()
	assert.Equal(t, nil, locales)

	os.Setenv("CS_API_KEY", oldApiKey)
}
