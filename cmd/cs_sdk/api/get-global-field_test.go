package api

import (
	"errors"
	"testing"

	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/stretchr/testify/assert"
)

func setupGetGlobalFieldTest() func() {
	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		if path == "global_fields/test" {
			return map[string]interface{}{}, nil

		}

		return nil, errors.New("invalid global field")
	}

	return func() {
		csSdkRequest = cs_sdk.Request
	}
}

func TestGetGlobalFieldSuccess(t *testing.T) {
	cleanup := setupGetGlobalFieldTest()
	defer cleanup()

	globalField := GetGlobalField("test")
	assert.Equal(t, map[string]interface{}{}, globalField)
}

func TestGetGlobalFieldErrInvalid(t *testing.T) {
	cleanup := setupGetGlobalFieldTest()
	defer cleanup()

	globalField := GetGlobalField("bogus")
	assert.Equal(t, nil, globalField)
}
