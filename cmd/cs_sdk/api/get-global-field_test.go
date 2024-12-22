package api

import (
	"testing"

	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/stretchr/testify/assert"
)

func TestGetGlobalField(t *testing.T) {
	init_env.Main("../../../.env.test")

	var contentType interface{}

	contentType = GetGlobalField("global_field")
	assert.NotEqual(t, nil, contentType)

	contentType = GetGlobalField("bogus")
	assert.Equal(t, nil, contentType)
}
