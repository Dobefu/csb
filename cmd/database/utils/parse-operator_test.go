package utils

import (
	"os"
	"testing"

	"github.com/Dobefu/csb/cmd/database/structs"
	"github.com/Dobefu/csb/cmd/logger"
	"github.com/stretchr/testify/assert"
)

func TestParseOperator(t *testing.T) {
	var operator string
	var err error

	logger.SetExitOnFatal(false)

	oldDbType := os.Getenv("DB_TYPE")
	dbTypes := []string{"mysql", "sqlite3", "postgres"}

	for _, dbType := range dbTypes {
		os.Setenv("DB_TYPE", dbType)

		operator, err = ParseOperator(structs.EQUALS)
		assert.Equal(t, operator, "=")
		assert.Equal(t, nil, err)

		operator, err = ParseOperator(structs.NOT_EQUALS)
		assert.Equal(t, operator, "<>")
		assert.Equal(t, nil, err)
	}

	operator, err = ParseOperator(9999)
	assert.Equal(t, operator, "")
	assert.NotEqual(t, nil, err)

	os.Setenv("DB_TYPE", "bogus")
	operator, err = ParseOperator(structs.EQUALS)
	assert.Equal(t, operator, "")
	assert.NotEqual(t, nil, err)

	os.Setenv("DB_TYPE", oldDbType)
}
