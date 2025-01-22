package api

import (
	"errors"
	"testing"

	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/stretchr/testify/assert"
)

func setupGetGlobalFieldsTest() func() {
	return func() {
		csSdkRequest = cs_sdk.Request
	}
}

func TestGetGlobalFieldsSuccess(t *testing.T) {
	cleanup := setupGetGlobalFieldsTest()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		return map[string]interface{}{}, nil
	}

	globalFields, err := GetGlobalFields()
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{}, globalFields)
}

func TestGetGlobalFieldsErr(t *testing.T) {
	cleanup := setupGetGlobalFieldsTest()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		return nil, errors.New("cannot get global fields")
	}

	globalFields, err := GetGlobalFields()
	assert.EqualError(t, err, "cannot get global fields")
	assert.Equal(t, map[string]interface{}(nil), globalFields)
}
