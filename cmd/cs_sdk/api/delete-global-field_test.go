package api

import (
	"errors"
	"testing"

	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/stretchr/testify/assert"
)

func setupDeleteGlobalFieldTest() func() {
	return func() {
		csSdkRequest = cs_sdk.Request
	}
}

func TestDeleteGlobalFieldSuccess(t *testing.T) {
	cleanup := setupDeleteGlobalFieldTest()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		return map[string]interface{}{}, nil
	}

	err := DeleteGlobalField("Test Name", false)
	assert.NoError(t, err)
}

func TestDeleteGlobalFieldErrDoesNotExist(t *testing.T) {
	cleanup := setupDeleteGlobalFieldTest()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		if method == "GET" {
			return nil, errors.New("cannot find global field")
		}

		return map[string]interface{}{}, nil
	}

	err := DeleteGlobalField("Bogus", false)
	assert.EqualError(t, err, "the global field does not exist")
}

func TestDeleteGlobalFieldErrDelete(t *testing.T) {
	cleanup := setupDeleteGlobalFieldTest()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		if method == "DELETE" {
			return nil, errors.New("cannot find global field")
		}

		return map[string]interface{}{}, nil
	}

	err := DeleteGlobalField("Test Name", false)
	assert.EqualError(t, err, "cannot find global field")
}
