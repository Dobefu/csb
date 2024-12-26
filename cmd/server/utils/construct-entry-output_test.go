package utils

import (
	"testing"

	api_structs "github.com/Dobefu/csb/cmd/api/structs"
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/stretchr/testify/assert"
)

func TestConstructEntryOutput(t *testing.T) {
	var breadcrumbs []structs.Route

	output := ConstructEntryOutput("test", []api_structs.AltLocale{}, breadcrumbs)
	assert.Equal(t, "test", output["data"]["entry"])
	assert.Equal(t, []api_structs.AltLocale{}, output["data"]["alt_locales"])
}
