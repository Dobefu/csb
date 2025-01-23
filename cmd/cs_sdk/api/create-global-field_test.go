package api

import (
	"errors"
	"testing"

	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/stretchr/testify/assert"
)

func setupCreateGlobalFieldTest() func() {
	return func() {
		csSdkRequest = cs_sdk.Request
	}
}

func TestCreateGlobalFieldSuccess(t *testing.T) {
	cleanup := setupCreateGlobalFieldTest()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		if method == "GET" {
			return nil, errors.New("cannot find global field")
		}

		return map[string]interface{}{}, nil
	}

	err := CreateGlobalField("Test", map[string]interface{}{})
	assert.NoError(t, err)
}

func TestCreateGlobalFieldErrAlreadyExists(t *testing.T) {
	cleanup := setupCreateGlobalFieldTest()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		return map[string]interface{}{}, nil
	}

	err := CreateGlobalField("Test", map[string]interface{}{})
	assert.EqualError(t, err, "the global field already exists")
}
