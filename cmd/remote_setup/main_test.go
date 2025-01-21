package remote_setup

import (
	"errors"
	"testing"

	"github.com/Dobefu/csb/cmd/cs_sdk/api"
	"github.com/stretchr/testify/assert"
)

func setupSetRouteTest() func() {
	apiCreateOrUpdateSeoGlobalField = func() error { return nil }
	apiCreateOrUpdateTranslationsContentType = func() error { return nil }

	return func() {
		apiCreateOrUpdateSeoGlobalField = api.CreateOrUpdateSeoGlobalField
		apiCreateOrUpdateTranslationsContentType = api.CreateOrUpdateTranslationsContentType
	}
}

func TestMainSuccess(t *testing.T) {
	cleanup := setupSetRouteTest()
	defer cleanup()

	err := Main()
	assert.NoError(t, err)
}

func TestMainErrCreateOrUpdateSeoGlobalField(t *testing.T) {
	cleanup := setupSetRouteTest()
	defer cleanup()

	apiCreateOrUpdateSeoGlobalField = func() error {
		return errors.New("cannot create or update SEO global field")
	}

	err := Main()
	assert.EqualError(t, err, "cannot create or update SEO global field")
}

func TestMainErrCreateOrUpdateTranslationsContentType(t *testing.T) {
	cleanup := setupSetRouteTest()
	defer cleanup()

	apiCreateOrUpdateTranslationsContentType = func() error {
		return errors.New("cannot create or update translations content type")
	}

	err := Main()
	assert.EqualError(t, err, "cannot create or update translations content type")
}
