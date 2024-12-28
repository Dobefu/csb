package api

import (
	"testing"

	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/stretchr/testify/assert"
)

func TestGetEntry(t *testing.T) {
	var entry map[string]interface{}
	var emptyEntry map[string]interface{}
	var err error

	entry, err = GetEntry(structs.Route{
		Uid:         "blt0617c28651fb44bf",
		ContentType: "basic_page",
		Locale:      "en",
	})
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, entry)

	entry, err = GetEntry(structs.Route{
		Uid:         "bogus",
		ContentType: "basic_page",
		Locale:      "en",
	})
	assert.NotEqual(t, nil, err)
	assert.Equal(t, emptyEntry, entry)
}
