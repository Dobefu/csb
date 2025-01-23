package api

import (
	"errors"
	"testing"

	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/stretchr/testify/assert"
)

func setupCreateOrUpdateTranslationsContentTypeTest() func() {
	return func() {
		csSdkRequest = cs_sdk.Request
	}
}

func TestCreateOrUpdateTranslationsContentTypeSuccess(t *testing.T) {
	cleanup := setupCreateOrUpdateTranslationsContentTypeTest()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		if method == "POST" {
			return nil, errors.New("cannot find content type")
		}

		return map[string]interface{}{}, nil
	}

	err := CreateOrUpdateTranslationsContentType()
	assert.NoError(t, err)
}

func TestCreateOrUpdateTranslationsContentTypeErrRequest(t *testing.T) {
	cleanup := setupCreateOrUpdateTranslationsContentTypeTest()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		if method != "PUT" {
			return map[string]interface{}{}, nil
		}

		return nil, errors.New("cannot create content type")
	}

	err := CreateOrUpdateTranslationsContentType()
	assert.EqualError(t, err, "cannot create content type")
}
