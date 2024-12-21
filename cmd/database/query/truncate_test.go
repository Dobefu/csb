package query

import (
	"os"
	"testing"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/Dobefu/csb/cmd/logger"
	"github.com/stretchr/testify/assert"
)

func TestTruncate(t *testing.T) {
	var err error

	init_env.Main("../../../.env.test")
	err = database.Connect()
	assert.Equal(t, nil, err)

	err = resetDb()
	assert.Equal(t, nil, err)

	err = Truncate("state")
	assert.Equal(t, nil, err)

	dbType := os.Getenv("DB_TYPE")
	os.Setenv("DB_TYPE", "bogus")
	logger.SetExitOnFatal(false)

	err = Truncate("state")
	assert.Equal(t, nil, err)

	os.Setenv("DB_TYPE", dbType)
}
