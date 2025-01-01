package cs_sdk

import (
	"os"
	"testing"

	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/stretchr/testify/assert"
)

func TestGetUrl(t *testing.T) {
	init_env.Main("../../.env.test")
	oldCsRegion := os.Getenv("CS_REGION")

	os.Setenv("CS_REGION", "eu")
	assert.Equal(t, GetUrl(true), "https://eu-api.contentstack.com")

	os.Setenv("CS_REGION", "us")
	assert.Equal(t, GetUrl(true), "https://api.contentstack.io")

	os.Setenv("CS_REGION", "azure-na")
	assert.Equal(t, GetUrl(true), "https://azure-na-api.contentstack.com")

	os.Setenv("CS_REGION", "azure-eu")
	assert.Equal(t, GetUrl(true), "https://azure-eu-api.contentstack.com")

	os.Setenv("CS_REGION", oldCsRegion)
}
