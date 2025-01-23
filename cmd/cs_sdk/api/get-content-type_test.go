package api

import (
	"errors"
	"testing"

	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/stretchr/testify/assert"
)

func setupGetContentTypeTest() func() {
	return func() {
		csSdkRequest = cs_sdk.Request
	}
}

func TestGetContentTypeSuccess(t *testing.T) {
	cleanup := setupGetContentTypeTest()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		return map[string]interface{}{}, nil
	}

	contentType := GetContentType("Test Name")
	assert.Equal(t, map[string]interface{}{}, contentType)
}

func TestGetContentTypeErrNotFound(t *testing.T) {
	cleanup := setupGetContentTypeTest()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		return nil, errors.New("cannot find content type")
	}

	contentType := GetContentType("Bogus")
	assert.Equal(t, nil, contentType)
}
