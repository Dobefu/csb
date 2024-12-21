package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructOutput(t *testing.T) {
	output := ConstructOutput()
	assert.NotEqual(t, nil, output)
}
