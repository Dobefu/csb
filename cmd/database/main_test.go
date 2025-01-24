package database

import (
	"database/sql"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupDatabaseTest() func() {
	sqlOpen = func(driverName, dataSourceName string) (*sql.DB, error) {
		if driverName == "bogus" {
			return nil, errors.New("invalid database type")
		}

		return nil, nil
	}

	return func() {
		sqlOpen = sql.Open

		os.Unsetenv("DB_CONN")
		os.Unsetenv("DB_TYPE")
	}
}

func TestConnectSuccess(t *testing.T) {
	cleanup := setupDatabaseTest()
	defer cleanup()

	os.Setenv("DB_CONN", "test-conn")
	os.Setenv("DB_TYPE", "test-type")

	err := Connect()
	assert.NoError(t, err)

}

func TestConnectErrInvalidDriver(t *testing.T) {
	cleanup := setupDatabaseTest()
	defer cleanup()

	os.Setenv("DB_CONN", "test-conn")
	os.Setenv("DB_TYPE", "bogus")

	err := Connect()
	assert.EqualError(t, err, "invalid database type")

}

func TestConnectErrNoConn(t *testing.T) {
	cleanup := setupDatabaseTest()
	defer cleanup()

	os.Setenv("DB_TYPE", "test-type")

	err := Connect()
	assert.EqualError(t, err, "DB_CONN is not set")

}

func TestConnectErrNoType(t *testing.T) {
	cleanup := setupDatabaseTest()
	defer cleanup()

	os.Setenv("DB_CONN", "test-conn")

	err := Connect()
	assert.EqualError(t, err, "DB_TYPE is not set")

}
