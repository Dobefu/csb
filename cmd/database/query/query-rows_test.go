package query

import (
	"database/sql"
	"os"
	"testing"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/database/structs"
	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/Dobefu/csb/cmd/logger"
	"github.com/stretchr/testify/assert"
)

func TestQueryRows(t *testing.T) {
	var row *sql.Rows
	var rowEmpty *sql.Rows
	var err error

	init_env.Main("../../../.env.test")
	err = database.Connect()
	assert.Equal(t, nil, err)

	err = resetDb()
	assert.Equal(t, nil, err)

	row, err = QueryRows("state", []string{"name"}, nil)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, rowEmpty, row)

	row, err = QueryRows("state", []string{"name"}, []structs.QueryWhere{
		{
			Name:     "name",
			Value:    "bogus",
			Operator: structs.EQUALS,
		},
	})
	assert.Equal(t, nil, err)
	assert.NotEqual(t, rowEmpty, row)

	dbType := os.Getenv("DB_TYPE")
	os.Setenv("DB_TYPE", "bogus")
	logger.SetExitOnFatal(false)

	row, err = QueryRows("state", []string{"name"}, nil)
	assert.Equal(t, nil, err)
	assert.Equal(t, rowEmpty, row)

	os.Setenv("DB_TYPE", dbType)
}
