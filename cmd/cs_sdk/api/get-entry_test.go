package api

import (
	"errors"
	"testing"

	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/stretchr/testify/assert"
)

func setupGetEntryTest() func() {
	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		if path == "content_types/test/entries/test?locale=en" {
			return map[string]interface{}{
				"entry": map[string]interface{}{},
			}, nil

		}

		return nil, errors.New("cannot get entry")
	}

	return func() {
		csSdkRequest = cs_sdk.Request
	}
}

func TestGetEntrySuccess(t *testing.T) {
	cleanup := setupGetEntryTest()
	defer cleanup()

	entry, err := GetEntry(structs.Route{
		Uid:         "test",
		ContentType: "test",
		Locale:      "en",
	})
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{
		"content_type": "test",
	}, entry)
}

func TestGetEntryErrInvalid(t *testing.T) {
	cleanup := setupGetEntryTest()
	defer cleanup()

	entry, err := GetEntry(structs.Route{Uid: "bogus", ContentType: "test"})
	assert.EqualError(t, err, "cannot get entry")
	assert.Equal(t, map[string]interface{}(nil), entry)
}
