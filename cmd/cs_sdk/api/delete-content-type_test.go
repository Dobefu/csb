package api

import (
	"errors"
	"testing"

	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/stretchr/testify/assert"
)

func setupDeleteContentTypeTest() func() {
	return func() {
		csSdkRequest = cs_sdk.Request
	}
}

func TestDeleteContentTypeSuccess(t *testing.T) {
	cleanup := setupDeleteContentTypeTest()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		return map[string]interface{}{}, nil
	}

	err := DeleteContentType("Test Name", false)
	assert.NoError(t, err)
}

func TestDeleteContentTypeErrDoesNotExist(t *testing.T) {
	cleanup := setupDeleteContentTypeTest()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		if method == "GET" {
			return nil, errors.New("cannot find content type")
		}

		return map[string]interface{}{}, nil
	}

	err := DeleteContentType("Bogus", false)
	assert.EqualError(t, err, "the content type does not exist")
}

func TestDeleteContentTypeErrDelete(t *testing.T) {
	cleanup := setupDeleteContentTypeTest()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		if method == "DELETE" {
			return nil, errors.New("cannot find content type")
		}

		return map[string]interface{}{}, nil
	}

	err := DeleteContentType("Test Name", false)
	assert.EqualError(t, err, "cannot find content type")
}
