package cs_sdk

import (
	"os"
	"testing"

	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	var data map[string]interface{}
	var emptyData map[string]interface{}
	var err error

	init_env.Main("../../.env.test")

	data, err = Request("content_types", "GET", nil, false)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, data)

	data, err = Request("content_types", "GET", map[string]interface{}{}, false)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, data)

	data, err = Request("content_types", "GET", nil, true)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, data)

	data, err = Request("content_types", "POST", nil, false)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, emptyData, data)

	oldCsRegion := os.Getenv("CS_REGION")
	os.Setenv("CS_REGION", "bogus")

	data, err = Request("content_types", "GET", nil, false)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, emptyData, data)

	os.Setenv("CS_REGION", oldCsRegion)
}
