package api

import (
	"errors"
	"testing"

	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/stretchr/testify/assert"
)

func setupGetContentTypesTest() func() {
	return func() {
		csSdkRequest = cs_sdk.Request
	}
}

func TestGetContentTypesSuccess(t *testing.T) {
	cleanup := setupGetContentTypesTest()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		return map[string]interface{}{}, nil
	}

	contentTypes, err := GetContentTypes()
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{}, contentTypes)
}

func TestGetContentTypesErr(t *testing.T) {
	cleanup := setupGetContentTypesTest()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		return nil, errors.New("cannot get content types")
	}

	contentTypes, err := GetContentTypes()
	assert.EqualError(t, err, "cannot get content types")
	assert.Equal(t, map[string]interface{}(nil), contentTypes)
}
