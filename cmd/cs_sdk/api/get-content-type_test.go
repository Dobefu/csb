package api

import (
	"testing"

	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/stretchr/testify/assert"
)

func TestGetContentType(t *testing.T) {
	init_env.Main("../../../.env.test")

	var contentType interface{}

	contentType = GetContentType("basic_page")
	assert.NotEqual(t, nil, contentType)

	contentType = GetContentType("bogus")
	assert.Equal(t, nil, contentType)
}
