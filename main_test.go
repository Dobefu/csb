package main

import (
	"errors"
	"flag"
	"os"
	"testing"

	"github.com/Dobefu/csb/cmd/check_health"
	"github.com/Dobefu/csb/cmd/cli"
	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/logger"
	"github.com/Dobefu/csb/cmd/migrate_db"
	"github.com/Dobefu/csb/cmd/remote_setup"
	"github.com/Dobefu/csb/cmd/remote_sync"
	"github.com/Dobefu/csb/cmd/server"
	"github.com/stretchr/testify/assert"
)

var err error

func TestInitDB(t *testing.T) {
	isLoggerFatalCalled := false
	mockLoggerFatal(&isLoggerFatalCalled)
	mockDatabaseConnectAndPing(nil, nil)

	defer cleanupLoggerFatal()
	defer cleanupDatabaseConnectAndPing()

	initDB()
	assert.False(t, isLoggerFatalCalled)
}

func TestInitDBErrPing(t *testing.T) {
	isLoggerFatalCalled := false
	mockLoggerFatal(&isLoggerFatalCalled)
	mockDatabaseConnectAndPing(nil, errors.New(""))

	defer cleanupLoggerFatal()
	defer cleanupDatabaseConnectAndPing()

	initDB()
	assert.True(t, isLoggerFatalCalled)
}

func TestInitDBErrConnect(t *testing.T) {
	isLoggerFatalCalled := false
	mockLoggerFatal(&isLoggerFatalCalled)
	mockDatabaseConnectAndPing(errors.New(""), errors.New(""))

	defer cleanupLoggerFatal()
	defer cleanupDatabaseConnectAndPing()

	initDB()
	assert.True(t, isLoggerFatalCalled)
}

func TestMainNoArguments(t *testing.T) {
	isExitCalled := false
	osExit = func(code int) { isExitCalled = true }
	mockDatabaseConnectAndPing(nil, nil)

	defer func() { osExit = os.Exit }()
	defer cleanupDatabaseConnectAndPing()

	main()
	assert.True(t, isExitCalled)
}

func TestMainWithSubCommand(t *testing.T) {
	isLoggerFatalCalled := false
	mockLoggerFatal(&isLoggerFatalCalled)
	oldOsArgs := os.Args
	mockDatabaseConnectAndPing(nil, nil)
	checkHealthMain = func() error { return nil }

	defer func() { os.Args = oldOsArgs }()
	defer cleanupLoggerFatal()
	defer cleanupDatabaseConnectAndPing()
	defer func() { checkHealthMain = check_health.Main }()

	os.Args = []string{os.Args[0], "check:health"}

	main()
	assert.False(t, isLoggerFatalCalled)
}

func TestMainWithSubCommandErr(t *testing.T) {
	isLoggerFatalCalled := false
	mockLoggerFatal(&isLoggerFatalCalled)
	oldOsArgs := os.Args
	mockDatabaseConnectAndPing(nil, nil)
	checkHealthMain = func() error { return errors.New("") }

	defer func() { os.Args = oldOsArgs }()
	defer cleanupLoggerFatal()
	defer cleanupDatabaseConnectAndPing()
	defer func() { checkHealthMain = check_health.Main }()

	os.Args = []string{os.Args[0], "check:health"}

	main()
	assert.True(t, isLoggerFatalCalled)
}

func TestRunSubCommandEmpty(t *testing.T) {
	isExitCalled := false
	osExit = func(code int) { isExitCalled = true }

	defer func() { osExit = os.Exit }()

	err = runSubCommand([]string{})
	assert.True(t, isExitCalled)
	assert.Equal(t, nil, err)
}

func TestRunSubCommandCheckHealth(t *testing.T) {
	mockDatabaseConnectAndPing(nil, nil)
	checkHealthMain = func() error { return nil }

	defer cleanupDatabaseConnectAndPing()
	defer func() { checkHealthMain = check_health.Main }()

	err = runSubCommand([]string{"check:health"})
	assert.Equal(t, nil, err)

	flagNewFlagSet = mockFlagNewFlagSet
	defer func() { flagNewFlagSet = flag.NewFlagSet }()

	err = runSubCommand([]string{"check:health", "--bogus"})
	assert.NotEqual(t, nil, err)
}

func TestRunSubCommandMigrateDb(t *testing.T) {
	mockDatabaseConnectAndPing(nil, nil)
	migrateDbMain = func(reset bool) error { return nil }

	defer cleanupDatabaseConnectAndPing()
	defer func() { migrateDbMain = migrate_db.Main }()

	err = runSubCommand([]string{"migrate:db"})
	assert.Equal(t, nil, err)

	flagNewFlagSet = mockFlagNewFlagSet
	defer func() { flagNewFlagSet = flag.NewFlagSet }()

	err = runSubCommand([]string{"migrate:db", "--bogus"})
	assert.NotEqual(t, nil, err)
}

func TestRunSubCommandRemoteSetup(t *testing.T) {
	mockDatabaseConnectAndPing(nil, nil)
	remoteSetupMain = func() error { return nil }

	defer cleanupDatabaseConnectAndPing()
	defer func() { remoteSetupMain = remote_setup.Main }()

	err = runSubCommand([]string{"remote:setup"})
	assert.Equal(t, nil, err)

	flagNewFlagSet = mockFlagNewFlagSet
	defer func() { flagNewFlagSet = flag.NewFlagSet }()

	err = runSubCommand([]string{"remote:setup", "--bogus"})
	assert.NotEqual(t, nil, err)
}

func TestRunSubCommandRemoteSync(t *testing.T) {
	mockDatabaseConnectAndPing(nil, nil)
	remoteSyncSync = func(reset bool) error { return nil }

	defer cleanupDatabaseConnectAndPing()
	defer func() { remoteSyncSync = remote_sync.Sync }()

	err = runSubCommand([]string{"remote:sync"})
	assert.Equal(t, nil, err)

	flagNewFlagSet = mockFlagNewFlagSet
	defer func() { flagNewFlagSet = flag.NewFlagSet }()

	err = runSubCommand([]string{"remote:sync", "--bogus"})
	assert.NotEqual(t, nil, err)
}

func TestRunSubCommandServer(t *testing.T) {
	mockDatabaseConnectAndPing(nil, nil)
	serverStart = func(port uint) error { return nil }

	defer cleanupDatabaseConnectAndPing()
	defer func() { serverStart = server.Start }()

	err = runSubCommand([]string{"server"})
	assert.Equal(t, nil, err)

	flagNewFlagSet = mockFlagNewFlagSet
	defer func() { flagNewFlagSet = flag.NewFlagSet }()

	err = runSubCommand([]string{"server", "--bogus"})
	assert.NotEqual(t, nil, err)
}

func TestRunSubCommandCreateContentType(t *testing.T) {
	mockDatabaseConnectAndPing(nil, nil)
	cliCreateContentType = func(isDryRun bool, name string, id string, withFields bool) error {
		return nil
	}

	defer cleanupDatabaseConnectAndPing()
	defer func() { cliCreateContentType = cli.CreateContentType }()

	err = runSubCommand([]string{"create:content-type"})
	assert.Equal(t, nil, err)

	flagNewFlagSet = mockFlagNewFlagSet
	defer func() { flagNewFlagSet = flag.NewFlagSet }()

	err = runSubCommand([]string{"create:content-type", "--bogus"})
	assert.NotEqual(t, nil, err)
}

func TestRunSubCommandFallback(t *testing.T) {
	isExitCalled := false
	osExit = func(code int) { isExitCalled = true }

	mockDatabaseConnectAndPing(nil, nil)

	defer func() { osExit = os.Exit }()
	defer cleanupDatabaseConnectAndPing()

	err = runSubCommand([]string{"bogus-cmd"})
	assert.Equal(t, nil, err)
	assert.True(t, isExitCalled)
}

func TestApplyGlobalFlagsVerbose(t *testing.T) {
	mockDatabaseConnectAndPing(nil, nil)

	oldVerbose := verbose
	*verbose = true

	defer func() { verbose = oldVerbose }()
	defer cleanupDatabaseConnectAndPing()

	applyGlobalFlags()
}

func TestApplyGlobalFlagsQuiet(t *testing.T) {
	mockDatabaseConnectAndPing(nil, nil)

	oldQuiet := quiet
	*quiet = true

	defer func() { quiet = oldQuiet }()
	defer cleanupDatabaseConnectAndPing()

	applyGlobalFlags()
}

func TestListSubCommands(t *testing.T) {
	isExitCalled := false
	osExit = func(code int) { isExitCalled = true }

	defer func() { osExit = os.Exit }()

	listSubCommands()
	assert.True(t, isExitCalled)
}

func mockFlagNewFlagSet(name string, errorHandling flag.ErrorHandling) *flag.FlagSet {
	return &flag.FlagSet{}
}

func mockLoggerFatal(isLoggerFatalCalled *bool) {
	loggerFatal = func(format string, a ...any) string {
		*isLoggerFatalCalled = true
		return ""
	}
}

func mockDatabaseConnectAndPing(connectErr, pingErr error) {
	databaseConnect = func() error { return connectErr }
	dbPing = func() error { return pingErr }
}

func cleanupLoggerFatal() {
	loggerFatal = logger.Fatal
}

func cleanupDatabaseConnectAndPing() {
	databaseConnect = database.Connect
	dbPing = database.DB.Ping
}
