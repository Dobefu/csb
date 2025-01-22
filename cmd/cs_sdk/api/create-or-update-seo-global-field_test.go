package api

import (
	"errors"
	"testing"

	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/stretchr/testify/assert"
)

func setupCreateOrUpdateSeoGlobalFieldTypeTest() func() {
	return func() {
		csSdkRequest = cs_sdk.Request
	}
}

func TestCreateOrUpdateSeoGlobalFieldTypeSuccess(t *testing.T) {
	cleanup := setupCreateOrUpdateSeoGlobalFieldTypeTest()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		if method == "POST" {
			return nil, errors.New("cannot find global field")
		}

		return map[string]interface{}{}, nil
	}

	err := CreateOrUpdateSeoGlobalField()
	assert.NoError(t, err)
}
