package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/Dobefu/csb/cmd/logger"
	"github.com/Dobefu/csb/cmd/migrate_db"
	"github.com/Dobefu/csb/cmd/remote_sync"
	"github.com/Dobefu/csb/cmd/server"
)

type subCommand struct {
	desc string
}

var (
	verbose = flag.Bool("verbose", false, "Enable verbose logging")
	quiet   = flag.Bool("quiet", false, "Only log warnings and errors")
	envPath = flag.String("env-file", ".env", "The location of the .env file. Defaults to .env")
)

func initDB() {
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
		err = flag.Parse(args[1:])

		if err != nil {
			break
		}

		applyGlobalFlags()

		err = migrate_db.Main(*reset)

	case "remote:sync":
		reset := flag.Bool("reset", false, "Synchronise all data, instead of starting from the last sync token")
		registerGlobalFlags(flag)
		err = flag.Parse(args[1:])

		if err != nil {
			break
		}

		applyGlobalFlags()

		err = remote_sync.Sync(*reset)

	case "server":
		port := flag.Uint("port", 4000, "The port to use for the web server")
		registerGlobalFlags(flag)
		err = flag.Parse(args[1:])

		if err != nil {
			break
		}

		applyGlobalFlags()
		err = server.Start(*port)

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

	init_env.Main(*envPath)
	initDB()
}

func listSubCommands() {
	cmds := map[string]subCommand{
		"migrate:db": {
			desc: "Migrate or initialise the database",
		},
		"remote:sync": {
			desc: "Synchronise Contentstack data into the database",
		},
		"server": {
			desc: "Run a webserver with API endpoints",
		},
	}

	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])

	for idx, cmd := range cmds {
		fmt.Printf("  %s:\n", idx)
		fmt.Printf("    %s\n", cmd.desc)
	}

	if flag.Lookup("test.v") == nil {
		os.Exit(1)
	}
}
