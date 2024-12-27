package api

import (
	"os"
	"testing"
	"time"

	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/stretchr/testify/assert"
)

func TestCreateContentType(t *testing.T) {
	var err error

	init_env.Main("../../../.env.test")

	ctName := "__csb_test"

	// If a test has failed before, clean up the content type now.
	_ = DeleteContentType(ctName, true)
	time.Sleep(time.Second / 2)

	err = CreateContentType(ctName, ctName, true)
	assert.Equal(t, nil, err)
	time.Sleep(time.Second / 2)

	err = CreateContentType(ctName, ctName, true)
	assert.NotEqual(t, nil, err)

	oldApiKey := os.Getenv("CS_API_KEY")
	os.Setenv("CS_API_KEY", "bogus")

	err = CreateContentType(ctName, ctName, true)
	assert.NotEqual(t, nil, err)

	os.Setenv("CS_API_KEY", oldApiKey)

	oldManagementToken := os.Getenv("CS_MANAGEMENT_TOKEN")
	os.Setenv("CS_MANAGEMENT_TOKEN", "bogus")

	err = DeleteContentType(ctName, false)
	assert.NotEqual(t, nil, err)

	os.Setenv("CS_MANAGEMENT_TOKEN", oldManagementToken)

	err = DeleteContentType(ctName, false)
	assert.Equal(t, nil, err)
	time.Sleep(time.Second / 2)

	err = DeleteContentType(ctName, false)
	assert.NotEqual(t, nil, err)
}
