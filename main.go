package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/migrate"
	"github.com/joho/godotenv"
)

type subCommand struct {
	flag *flag.FlagSet
	desc string
}

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("No .env file found. Please copy it from the .env.example and enter your credentials")
	}

	err = database.Connect()

	if err != nil {
		log.Fatalln("Could not connect to the database: " + err.Error())
	}

	err = database.DB.Ping()

	if err != nil {
		log.Fatalln("Could not connect to the database: " + err.Error())
	}
}

func main() {
	cmdName, cmd := parseSubCommands()
	var err error

	switch cmdName {
	case "migrate":
		reset := cmd.flag.Bool("reset", false, "Migrate from a clean database. Warning: this will delete existing data")
		cmd.flag.Parse(os.Args[2:])

		err = migrate.Main(*reset)
		break
	default:
		break
	}

	if err != nil {
		log.Fatalln(err)
	}
}

func parseSubCommands() (string, subCommand) {
	cmds := map[string]subCommand{
		"migrate": {
			flag: flag.NewFlagSet("migrate", flag.ExitOnError),
			desc: "Migrate or initialise the database",
		},
	}

	if len(os.Args) < 2 {
		listCmds(cmds)
	}

	subCmd, subCmdExists := cmds[os.Args[1]]

	if !subCmdExists {
		listCmds(cmds)
	}

	return os.Args[1], subCmd
}

func listCmds(cmds map[string]subCommand) {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])

	for idx, cmd := range cmds {
		fmt.Printf("  %s:\n", idx)
		fmt.Printf("    %s\n", cmd.desc)
	}

	os.Exit(1)
}
