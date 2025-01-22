package migrate_db

import (
	"database/sql"
	"errors"
	"testing"
	"testing/fstest"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Dobefu/csb/cmd/database/query"
	"github.com/Dobefu/csb/cmd/database/structs"
	"github.com/stretchr/testify/assert"
)

func setupMigrateDbTest(t *testing.T) (sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	queryQueryRaw = func(sql string) (sql.Result, error) {
		return nil, nil
	}

	queryQueryRow = func(table string, fields []string, where []structs.QueryWhere) *sql.Row {
		return db.QueryRow("SELECT (.+) FROM migrations", nil)
	}

	queryTruncate = func(table string) error {
		return nil
	}

	queryInsert = func(table string, values []structs.QueryValue) error {
		return nil
	}

	cleanup := func() {
		queryQueryRaw = query.QueryRaw
		queryQueryRow = query.QueryRow
		queryTruncate = query.Truncate
		queryInsert = query.Insert
		getFs = func() FS { return content }
	}

	return mock, cleanup
}

func TestMainSuccess(t *testing.T) {
	mock, cleanup := setupMigrateDbTest(t)
	defer cleanup()

	cols := []string{"version", "dirty"}
	mock.ExpectQuery("SELECT (.+) FROM migrations").WillReturnRows(
		sqlmock.NewRows(cols).AddRow(1, false),
	)

	err := Main(true)
	assert.NoError(t, err)
}

func TestMainErrDown(t *testing.T) {
	_, cleanup := setupMigrateDbTest(t)
	defer cleanup()

	queryQueryRaw = func(sql string) (sql.Result, error) {
		return nil, errors.New("QueryRaw failed")
	}

	err := Main(true)
	assert.EqualError(t, err, "QueryRaw failed")
}

func TestMainErrUp(t *testing.T) {
	_, cleanup := setupMigrateDbTest(t)
	defer cleanup()

	queryQueryRaw = func(sql string) (sql.Result, error) {
		return nil, errors.New("QueryRaw failed")
	}

	err := Main(false)
	assert.EqualError(t, err, "QueryRaw failed")
}

func TestDownNothingToMigrate(t *testing.T) {
	_, cleanup := setupMigrateDbTest(t)
	defer cleanup()

	err := down()
	assert.NoError(t, err)
}

func TestDownErrReadDir(t *testing.T) {
	mock, cleanup := setupMigrateDbTest(t)
	defer cleanup()

	getFs = func() FS {
		return fstest.MapFS{
			"bogus/000001-test.up.sql": {Data: []byte("")},
		}
	}

	cols := []string{"version", "dirty"}
	mock.ExpectQuery("SELECT (.+) FROM migrations").WillReturnRows(
		sqlmock.NewRows(cols).AddRow(1, false),
	)

	err := down()
	assert.EqualError(t, err, "open migrations: file does not exist")
}

func TestDownErrReadFile(t *testing.T) {
	mock, cleanup := setupMigrateDbTest(t)
	defer cleanup()

	getFs = func() FS {
		return fstest.MapFS{
			"migrations/000001-test.down.sql/": {Data: []byte("")},
		}
	}

	cols := []string{"version", "dirty"}
	mock.ExpectQuery("SELECT (.+) FROM migrations").WillReturnRows(
		sqlmock.NewRows(cols).AddRow(1, false),
	)

	err := down()
	assert.EqualError(t, err, "read migrations/000001-test.down.sql: invalid argument")
}

func TestDownErrSetMigrationState(t *testing.T) {
	mock, cleanup := setupMigrateDbTest(t)
	defer cleanup()

	queryTruncate = func(table string) error {
		return errors.New("could not truncate table")
	}

	cols := []string{"version", "dirty"}
	mock.ExpectQuery("SELECT (.+) FROM migrations").WillReturnRows(
		sqlmock.NewRows(cols).AddRow(1, false),
	)

	err := down()
	assert.EqualError(t, err, "could not truncate table")
}

func TestUpErrReadDir(t *testing.T) {
	_, cleanup := setupMigrateDbTest(t)
	defer cleanup()

	getFs = func() FS {
		return fstest.MapFS{
			"bogus/000001-test.up.sql": {Data: []byte("")},
		}
	}

	err := up()
	assert.EqualError(t, err, "open migrations: file does not exist")
}

func TestUpErrRunMigration(t *testing.T) {
	_, cleanup := setupMigrateDbTest(t)
	defer cleanup()

	getFs = func() FS {
		return fstest.MapFS{
			"migrations/000001-test.up.sql/": {Data: []byte("")},
		}
	}

	err := up()
	assert.EqualError(t, err, "read migrations/000001-test.up.sql: invalid argument")
}

func TestUpFromIndex(t *testing.T) {
	mock, cleanup := setupMigrateDbTest(t)
	defer cleanup()

	getFs = func() FS {
		return fstest.MapFS{
			"migrations/000001-test.up.sql": {Data: []byte("")},
			"migrations/000002-test.up.sql": {Data: []byte("")},
		}
	}

	cols := []string{"version", "dirty"}
	mock.ExpectQuery("SELECT (.+) FROM migrations").WillReturnRows(
		sqlmock.NewRows(cols).AddRow(1, false),
	)

	err := up()
	assert.NoError(t, err)
}

func TestUpErrSetMigrationState(t *testing.T) {
	_, cleanup := setupMigrateDbTest(t)
	defer cleanup()

	queryInsert = func(table string, values []structs.QueryValue) error {
		return errors.New("could not insert row")
	}

	err := up()
	assert.EqualError(t, err, "could not insert row")
}

func TestSetMigrationStateErrCreateMigrationsTable(t *testing.T) {
	_, cleanup := setupMigrateDbTest(t)
	defer cleanup()

	queryQueryRaw = func(sql string) (sql.Result, error) {
		return nil, errors.New("QueryRaw failed")
	}

	err := setMigrationState(1, false)
	assert.EqualError(t, err, "QueryRaw failed")
}

func TestRunMigrationErrQueryRaw(t *testing.T) {
	_, cleanup := setupMigrateDbTest(t)
	defer cleanup()

	getFs = func() FS {
		return fstest.MapFS{
			"migrations/000001-test.up.sql": {Data: []byte("")},
		}
	}

	queryQueryRaw = func(sql string) (sql.Result, error) {
		return nil, errors.New("QueryRaw failed")
	}

	err := runMigration("000001-test.up.sql", 1)
	assert.EqualError(t, err, "QueryRaw failed")
}
