package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateId(t *testing.T) {
	assert.Equal(t, "testen", GenerateId("test", "en"))
	assert.Equal(t, "testnl-nl", GenerateId("test", "nl-nl"))
	assert.Equal(t, "test", GenerateId("test", ""))
	assert.Equal(t, "en", GenerateId("", "en"))
}
