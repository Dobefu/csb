package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	err := Start(40000)
	assert.Equal(t, nil, err)
}
