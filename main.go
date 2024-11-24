package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/logger"
	"github.com/Dobefu/csb/cmd/migrate_db"
	"github.com/Dobefu/csb/cmd/remote_sync"

	_ "github.com/Dobefu/csb/cmd/init"
)

type subCommand struct {
	flag *flag.FlagSet
	desc string
}

func init() {
	err := database.Connect()

	if err != nil {
		logger.Fatal("Could not connect to the database: %s", err.Error())
	}

	err = database.DB.Ping()

	if err != nil {
		logger.Fatal("Could not connect to the database: %s", err.Error())
	}
}

func main() {
	cmdName, cmd := parseSubCommands()
	var err error

	switch cmdName {
	case "migrate:db":
		reset := cmd.flag.Bool("reset", false, "Migrate from a clean database. Warning: this will delete existing data")
		verbose := cmd.flag.Bool("verbose", false, "Enable verbose logging")
		cmd.flag.Parse(os.Args[2:])

		if *verbose {
			logger.SetLogLevel(logger.LOG_VERBOSE)
		}

		err = migrate_db.Main(*reset)
		break

	case "remote:sync":
		reset := cmd.flag.Bool("reset", false, "Synchronise all data, instead of starting from the last sync token")
		verbose := cmd.flag.Bool("verbose", false, "Enable verbose logging")
		cmd.flag.Parse(os.Args[2:])

		if *verbose {
			logger.SetLogLevel(logger.LOG_VERBOSE)
		}

		err = remote_sync.Sync(*reset)
		break
	default:
		break
	}

	if err != nil {
		logger.Fatal(err.Error())
	}
}

func getSubCommands() map[string]subCommand {
	return map[string]subCommand{
		"migrate:db": {
			desc: "Migrate or initialise the database",
		},
		"remote:sync": {
			desc: "Synchronise Contentstack data into the database",
		},
	}

}

func parseSubCommands() (string, subCommand) {
	cmds := getSubCommands()

	for cmdName, cmd := range cmds {
		cmds[cmdName] = subCommand{
			flag: flag.NewFlagSet(cmdName, flag.ExitOnError),
			desc: cmd.desc,
		}
	}

	if len(os.Args) < 2 {
		listCmds()
	}

	subCmd, subCmdExists := cmds[os.Args[1]]

	if !subCmdExists {
		listCmds()
	}

	return os.Args[1], subCmd
}

func listCmds() {
	cmds := getSubCommands()

	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])

	for idx, cmd := range cmds {
		fmt.Printf("  %s:\n", idx)
		fmt.Printf("    %s\n", cmd.desc)
	}

	os.Exit(1)
}
