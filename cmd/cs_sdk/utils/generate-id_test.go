package utils

import (
	"testing"

	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/stretchr/testify/assert"
)

func TestGenerateId(t *testing.T) {
	assert.Equal(t, "testen", GenerateId(structs.Route{Uid: "test", Locale: "en"}))
	assert.Equal(t, "testnl-nl", GenerateId(structs.Route{Uid: "test", Locale: "nl-nl"}))
	assert.Equal(t, "test", GenerateId(structs.Route{Uid: "test"}))
	assert.Equal(t, "en", GenerateId(structs.Route{Locale: "en"}))
}
