package api

import (
	"errors"
	"testing"

	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/stretchr/testify/assert"
)

func setupCreateContentTypeTest() func() {
	return func() {
		csSdkRequest = cs_sdk.Request
	}
}

func TestCreateContentTypeSuccess(t *testing.T) {
	cleanup := setupCreateContentTypeTest()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		if method == "GET" {
			return nil, errors.New("cannot find content type")
		}

		return map[string]interface{}{}, nil
	}

	err := CreateContentType("Test Name", "test_id", true)
	assert.NoError(t, err)
}

func TestCreateContentTypeErr(t *testing.T) {
	cleanup := setupCreateContentTypeTest()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		if method == "GET" {
			return nil, errors.New("cannot find content type")
		}

		if path == "global_fields" || path == "global_fields/seo" {
			return nil, errors.New("cannot create or update global field")
		}

		return map[string]interface{}{}, nil
	}

	err := CreateContentType("Test Name", "test_id", true)
	assert.EqualError(t, err, "cannot create or update global field")
}
