package utils

import (
	"os"
	"testing"

	"github.com/Dobefu/csb/cmd/database/structs"
	"github.com/stretchr/testify/assert"
)

func TestParseOperator(t *testing.T) {
	var operator string
	var err error

	dbTypes := []string{"mysql", "sqlite3", "postgres"}

	for _, dbType := range dbTypes {
		os.Setenv("DB_TYPE", dbType)

		operator, err = ParseOperator(structs.EQUALS)
		assert.Equal(t, operator, "=")
		assert.Equal(t, err, nil)
		operator, err = ParseOperator(structs.NOT_EQUALS)
		assert.Equal(t, operator, "<>")
		assert.Equal(t, err, nil)
	}

	operator, err = ParseOperator(9999)
	assert.Equal(t, operator, "")
	assert.NotEqual(t, err, nil)
}
