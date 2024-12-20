package api

import (
	"testing"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/stretchr/testify/assert"
)

func TestGetContentTypes(t *testing.T) {
	_, err := GetContentTypes()
	assert.NotEqual(t, nil, err)

	init_env.Main("../../.env.test")

	err = database.Connect()
	assert.Equal(t, nil, err)

	contentTypes, err := GetContentTypes()
	assert.Equal(t, nil, err)

	assert.NotEqual(t, nil, contentTypes)
}
