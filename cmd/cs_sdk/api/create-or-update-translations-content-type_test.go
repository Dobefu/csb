package api

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateOrUpdateTranslationsContentType(t *testing.T) {
	var err error

	oldApiKey := os.Getenv("CS_API_KEY")
	os.Setenv("CS_API_KEY", "bogus")

	err = CreateOrUpdateTranslationsContentType()
	assert.NotEqual(t, nil, err)

	os.Setenv("CS_API_KEY", oldApiKey)

	err = CreateOrUpdateTranslationsContentType()
	assert.Equal(t, nil, err)
}
