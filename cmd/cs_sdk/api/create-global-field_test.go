package api

import (
	"os"
	"testing"
	"time"

	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/stretchr/testify/assert"
)

func TestCreateGlobalField(t *testing.T) {
	var err error

	init_env.Main("../../../.env.test")

	gfName := "__csb_test"
	gfData := map[string]interface{}{
		"global_field": map[string]interface{}{
			"title": gfName,
			"uid":   gfName,
			"schema": []map[string]interface{}{
				{
					"display_name": "Title",
					"uid":          "title",
					"data_type":    "text",
				},
			},
		},
	}

	// If a test has failed before, clean up the global field now.
	_ = DeleteGlobalField(gfName, true)
	time.Sleep(time.Second / 2)

	err = CreateGlobalField(gfName, gfData)
	assert.Equal(t, nil, err)
	time.Sleep(time.Second / 2)

	err = CreateGlobalField(gfName, gfData)
	assert.NotEqual(t, nil, err)

	oldApiKey := os.Getenv("CS_API_KEY")
	os.Setenv("CS_API_KEY", "bogus")

	err = CreateGlobalField(gfName, gfData)
	assert.NotEqual(t, nil, err)

	os.Setenv("CS_API_KEY", oldApiKey)

	oldManagementToken := os.Getenv("CS_MANAGEMENT_TOKEN")
	os.Setenv("CS_MANAGEMENT_TOKEN", "bogus")

	err = DeleteGlobalField(gfName, false)
	assert.NotEqual(t, nil, err)

	os.Setenv("CS_MANAGEMENT_TOKEN", oldManagementToken)

	err = DeleteGlobalField(gfName, false)
	assert.Equal(t, nil, err)
	time.Sleep(time.Second / 2)

	err = DeleteGlobalField(gfName, false)
	assert.NotEqual(t, nil, err)
}
