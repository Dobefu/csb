package cs_sdk

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogMessage(t *testing.T) {
	os.Setenv("CS_REGION", "eu")
	assert.Equal(t, GetUrl(true), "https://eu-api.contentstack.com")

	os.Setenv("CS_REGION", "us")
	assert.Equal(t, GetUrl(true), "https://api.contentstack.io")

	os.Setenv("CS_REGION", "azure-na")
	assert.Equal(t, GetUrl(true), "https://azure-na-api.contentstack.com")

	os.Setenv("CS_REGION", "azure-eu")
	assert.Equal(t, GetUrl(true), "https://azure-eu-api.contentstack.com")
}
