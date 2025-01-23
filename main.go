package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Dobefu/csb/cmd/color"
	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/Dobefu/csb/cmd/logger"
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
	err := databaseConnect()

	if err != nil {
		loggerFatal("Could not connect to the database: %s", err.Error())
	}

	err = dbPing()

	if err != nil {
		loggerFatal("Could not ping the database: %s", err.Error())
	}
}

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		listSubCommands()
		return
	}

	err := runSubCommand(args)

	if err != nil {
		loggerFatal(err.Error())
	}
}

func runSubCommand(args []string) error {
	if len(args) < 1 {
		listSubCommands()
		return nil
	}

	flag := flagNewFlagSet(args[0], flag.ExitOnError)
	var err error

	switch args[0] {
	case "check:health":
		registerGlobalFlags(flag)
		err = flag.Parse(args[1:])

		if err != nil {
			break
		}

		applyGlobalFlags()

		err = checkHealthMain()

	case "migrate:db":
		reset := flag.Bool("reset", false, "Migrate from a clean database. Warning: this will delete existing data")

		registerGlobalFlags(flag)
		err = flag.Parse(args[1:])

		if err != nil {
			break
		}

		applyGlobalFlags()

		err = migrateDbMain(*reset)

	case "remote:setup":
		registerGlobalFlags(flag)
		err = flag.Parse(args[1:])

		if err != nil {
			break
		}

		applyGlobalFlags()

		err = remoteSetupMain()

	case "remote:sync":
		reset := flag.Bool("reset", false, "Synchronise all data, instead of starting from the last sync token")

		registerGlobalFlags(flag)
		err = flag.Parse(args[1:])

		if err != nil {
			break
		}

		applyGlobalFlags()

		err = remoteSyncSync(*reset)

	case "server":
		port := flag.Uint("port", 4000, "The port to use for the web server")

		registerGlobalFlags(flag)
		err = flag.Parse(args[1:])

		if err != nil {
			break
		}

		applyGlobalFlags()
		err = serverStart(*port)

	case "create:content-type":
		name := flag.String("name", "", "The name of the content type to create")
		machineName := flag.String("machine-name", "", "The machine name of the content type to create")
		isDryRun := flag.Bool("dry-run", false, "Log the actions instead of actually running them")

		registerGlobalFlags(flag)
		err = flag.Parse(args[1:])

		if err != nil {
			break
		}

		applyGlobalFlags()
		err = cliCreateContentType(*isDryRun, *name, *machineName, true)

	default:
		listSubCommands()
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
		"check:health": {
			desc: "Validate the health of the application configuration",
		},
		"migrate:db": {
			desc: "Migrate or initialise the database",
		},
		"remote:setup": {
			desc: "Set up or update necessary config in Contentstack",
		},
		"remote:sync": {
			desc: "Synchronise all Contentstack entries into the database",
		},
		"server": {
			desc: "Run a webserver with API endpoints",
		},
		"create:content-type": {
			desc: "Create a new content type",
		},
	}

	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])

	for idx, cmd := range cmds {
		color.PrintfColor(color.FgLightCyan, color.BgDefault, "  %s:\n", idx)
		fmt.Printf("    %s\n", cmd.desc)
	}

	osExit(1)
}
