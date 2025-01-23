package main

import (
	"flag"
	"os"

	"github.com/Dobefu/csb/cmd/check_health"
	"github.com/Dobefu/csb/cmd/cli"
	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/logger"
	"github.com/Dobefu/csb/cmd/migrate_db"
	"github.com/Dobefu/csb/cmd/remote_setup"
	"github.com/Dobefu/csb/cmd/remote_sync"
	"github.com/Dobefu/csb/cmd/server"
)

var databaseConnect = database.Connect
var dbPing = func() error { return database.DB.Ping() }
var flagNewFlagSet = flag.NewFlagSet

var checkHealthMain = check_health.Main
var migrateDbMain = migrate_db.Main
var remoteSetupMain = remote_setup.Main
var remoteSyncSync = remote_sync.Sync
var serverStart = server.Start
var cliCreateContentType = cli.CreateContentType

var loggerFatal = logger.Fatal
var osExit = os.Exit
