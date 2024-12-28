package remote_setup

import (
	"testing"

	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	var err error

	err = Main()
	assert.NotEqual(t, nil, err)

	init_env.Main("../../.env.test")

	err = Main()
	assert.Equal(t, nil, err)
}
