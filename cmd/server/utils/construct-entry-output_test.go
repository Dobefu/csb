package utils

import (
	"testing"

	"github.com/Dobefu/csb/cmd/api/structs"
	"github.com/stretchr/testify/assert"
)

func TestConstructEntryOutput(t *testing.T) {
	output := ConstructEntryOutput("test", []structs.AltLocale{})
	assert.Equal(t, "test", output["data"]["entry"])
	assert.Equal(t, []structs.AltLocale{}, output["data"]["alt_locales"])
}
