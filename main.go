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
	parseGlobalFlags()
	args := flag.Args()

	cmdName, cmd := parseSubCommands()
	var err error

	switch cmdName {
	case "migrate:db":
		reset := cmd.flag.Bool("reset", false, "Migrate from a clean database. Warning: this will delete existing data")
		cmd.flag.Parse(args[1:])

		err = migrate_db.Main(*reset)
		break

	case "remote:sync":
		reset := cmd.flag.Bool("reset", false, "Synchronise all data, instead of starting from the last sync token")
		cmd.flag.Parse(args[1:])

		err = remote_sync.Sync(*reset)
		break
	default:
		break
	}

	if err != nil {
		logger.Fatal(err.Error())
	}
}

func parseGlobalFlags() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")

	if *verbose {
		logger.SetLogLevel(logger.LOG_VERBOSE)
	}

	quiet := flag.Bool("quiet", false, "Only log warnings and errors")

	if *quiet {
		logger.SetLogLevel(logger.LOG_WARNING)
	}

	flag.Parse()
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
	args := flag.Args()

	for cmdName, cmd := range cmds {
		cmds[cmdName] = subCommand{
			flag: flag.NewFlagSet(cmdName, flag.ExitOnError),
			desc: cmd.desc,
		}
	}

	if len(args) < 1 {
		flag.Usage()
		listSubCommands()
	}

	subCmd, subCmdExists := cmds[args[0]]

	if !subCmdExists {
		listSubCommands()
	}

	return args[0], subCmd
}

func listSubCommands() {
	cmds := getSubCommands()

	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])

	for idx, cmd := range cmds {
		fmt.Printf("  %s:\n", idx)
		fmt.Printf("    %s\n", cmd.desc)
	}

	if flag.Lookup("test.v") == nil {
		os.Exit(1)
	}
}
