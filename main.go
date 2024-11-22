package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Dobefu/csb/cmd/database"
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
	parseSubCommands()
}

func parseSubCommands() {
	cmds := map[string]subCommand{
		"init": {
			flag: flag.NewFlagSet("Init", flag.ExitOnError),
			desc: "Initialise the database",
		},
	}

	if len(os.Args) < 2 {
		listCmds(cmds)
	}

	_, subCmdExists := cmds[os.Args[1]]

	if !subCmdExists {
		listCmds(cmds)
	}

	flag.PrintDefaults()
}

func listCmds(cmds map[string]subCommand) {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	fmt.Println("")

	for idx, cmd := range cmds {
		fmt.Printf("%s:\n", idx)
		fmt.Printf("  %s\n", cmd.desc)
	}

	fmt.Println("")
	os.Exit(1)
}
