package api

import (
	"errors"
	"testing"

	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/stretchr/testify/assert"
)

func setupGetContentTypeTest() func() {
	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		if path == "content_types/test" {
			return map[string]interface{}{}, nil

		}

		return nil, errors.New("invalid content type")
	}

	return func() {
		csSdkRequest = cs_sdk.Request
	}
}

func TestGetContentTypeSuccess(t *testing.T) {
	cleanup := setupGetContentTypeTest()
	defer cleanup()

	contentType, err := GetContentType("test")
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{}, contentType)
}

func TestGetContentTypeErrInvalid(t *testing.T) {
	cleanup := setupGetContentTypeTest()
	defer cleanup()

	contentType, err := GetContentType("bogus")
	assert.EqualError(t, err, "invalid content type")
	assert.Equal(t, map[string]interface{}(nil), contentType)
}
