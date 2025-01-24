package cs_sdk

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupGetUrlTest() func() {
	return func() {
		os.Unsetenv("CS_REGION")
	}
}

func TestGetUrlEu(t *testing.T) {
	cleanup := setupGetUrlTest()
	defer cleanup()

	os.Setenv("CS_REGION", "eu")
	assert.Equal(t, GetUrl(true), "https://eu-api.contentstack.com")
}

func TestGetUrlUs(t *testing.T) {
	cleanup := setupGetUrlTest()
	defer cleanup()

	os.Setenv("CS_REGION", "us")
	assert.Equal(t, GetUrl(true), "https://api.contentstack.io")
}

func TestGetUrlAzureNa(t *testing.T) {
	cleanup := setupGetUrlTest()
	defer cleanup()

	os.Setenv("CS_REGION", "azure-na")
	assert.Equal(t, GetUrl(true), "https://azure-na-api.contentstack.com")
}

func TestGetUrlAzureEu(t *testing.T) {
	cleanup := setupGetUrlTest()
	defer cleanup()

	os.Setenv("CS_REGION", "azure-eu")
	assert.Equal(t, GetUrl(true), "https://azure-eu-api.contentstack.com")
}
