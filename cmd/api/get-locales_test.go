package api

import (
	"errors"
	"testing"

	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/stretchr/testify/assert"
)

func setupGetLocalesTest() func() {
	return func() {
		csSdkRequest = cs_sdk.Request
	}
}

func TestGetLocalesSuccess(t *testing.T) {
	cleanup := setupGetLocalesTest()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		return map[string]interface{}{}, nil
	}

	locales, err := GetLocales()
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{}, locales)
}

func TestGetLocalesErr(t *testing.T) {
	cleanup := setupGetLocalesTest()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		return nil, errors.New("cannot get locales")
	}

	locales, err := GetLocales()
	assert.EqualError(t, err, "cannot get locales")
	assert.Equal(t, map[string]interface{}(nil), locales)
}
