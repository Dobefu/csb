package check_health

import (
	"os"
	"testing"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	var err error

	err = Main()
	assert.NotEqual(t, nil, err)

	init_env.Main("../../.env.test")
	err = database.Connect()
	assert.Equal(t, nil, err)

	err = Main()
	assert.Equal(t, nil, err)

	oldApiKey := os.Getenv("CS_API_KEY")
	os.Setenv("CS_API_KEY", "bogus")

	err = Main()
	assert.NotEqual(t, nil, err)

	os.Setenv("CS_API_KEY", oldApiKey)
}
