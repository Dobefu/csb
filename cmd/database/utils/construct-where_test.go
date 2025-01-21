package utils

import (
	"os"
	"testing"

	"github.com/Dobefu/csb/cmd/database/structs"
	"github.com/Dobefu/csb/cmd/logger"
	"github.com/stretchr/testify/assert"
)

func setupConstructWhereTest() func() {
	env := os.Getenv("DB_TYPE")
	os.Setenv("DB_TYPE", "mysql")
	logger.SetExitOnFatal(false)

	return func() {
		os.Setenv("DB_TYPE", env)
		logger.SetExitOnFatal(true)
	}
}

func TestConstructWhereSuccess(t *testing.T) {
	cleanup := setupConstructWhereTest()
	defer cleanup()

	where, args := ConstructWhere([]structs.QueryWhere{
		{
			Name:  "name",
			Value: "value",
		},
		{
			Name:  "name2",
			Value: "value2",
		},
	})
	assert.Equal(t, "WHERE name = ? AND name2 = ?", where)
	assert.Equal(t, []interface{}{"value", "value2"}, args)
}

func TestConstructWhereErrParseOperator(t *testing.T) {
	cleanup := setupConstructWhereTest()
	defer cleanup()

	os.Setenv("DB_TYPE", "bogus")

	where, args := ConstructWhere([]structs.QueryWhere{
		{
			Name:  "name",
			Value: "value",
		},
	})
	assert.Equal(t, "WHERE ", where)
	assert.Equal(t, []interface{}(nil), args)
}
