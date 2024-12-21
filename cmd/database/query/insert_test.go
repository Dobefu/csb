package query

import (
	"os"
	"testing"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/database/structs"
	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/Dobefu/csb/cmd/logger"
	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	var err error

	init_env.Main("../../../.env.test")
	err = database.Connect()
	assert.Equal(t, nil, err)

	_, err = QueryRaw("DROP TABLE IF EXISTS state;")
	assert.Equal(t, nil, err)

	_, err = QueryRaw(`CREATE TABLE IF NOT EXISTS state(
  name varchar(255) NOT NULL PRIMARY KEY UNIQUE,
  value varchar(255) NOT NULL
);
  `)
	assert.Equal(t, nil, err)

	err = Insert("state", []structs.QueryValue{
		{
			Name:  "name",
			Value: "test",
		},
		{
			Name:  "value",
			Value: "test",
		},
	})
	assert.Equal(t, nil, err)

	dbType := os.Getenv("DB_TYPE")
	os.Setenv("DB_TYPE", "")
	logger.SetExitOnFatal(false)

	err = Insert("state", []structs.QueryValue{})
	assert.Equal(t, nil, err)

	os.Setenv("DB_TYPE", dbType)
}
