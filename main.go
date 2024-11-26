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

var (
	verbose = flag.Bool("verbose", false, "Enable verbose logging")
	quiet   = flag.Bool("quiet", false, "Only log warnings and errors")
)

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
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		listSubCommands()
	}

	err := runSubCommand(args)

	if err != nil {
		logger.Fatal(err.Error())
	}
}

func runSubCommand(args []string) error {
	flag := flag.NewFlagSet(args[0], flag.ExitOnError)
	var err error

	switch args[0] {
	case "migrate:db":
		reset := flag.Bool("reset", false, "Migrate from a clean database. Warning: this will delete existing data")
		registerGlobalFlags(flag)
		flag.Parse(args[1:])
		applyGlobalFlags()

		err = migrate_db.Main(*reset)
		break

	case "remote:sync":
		reset := flag.Bool("reset", false, "Synchronise all data, instead of starting from the last sync token")
		registerGlobalFlags(flag)
		flag.Parse(args[1:])
		applyGlobalFlags()

		err = remote_sync.Sync(*reset)
		break
	default:
		break
	}

	return err
}

func registerGlobalFlags(fset *flag.FlagSet) {
	flag.VisitAll(func(f *flag.Flag) {
		fset.Var(f.Value, f.Name, f.Usage)
	})
}

func applyGlobalFlags() {
	if *verbose {
		logger.SetLogLevel(logger.LOG_VERBOSE)
	}

	if *quiet {
		logger.SetLogLevel(logger.LOG_WARNING)
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
