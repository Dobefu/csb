package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	setArgs("migrate:db", "--reset", "--verbose", "--env-file=.env.test")
	main()

	setArgs("remote:sync", "--reset", "--quiet", "--env-file=.env.test")
	main()
}

func setArgs(args ...string) {
	os.Args = []string{os.Args[0]}
	os.Args = append(os.Args, args...)
}

func TestListSubcommands(t *testing.T) {
	listSubCommands()
}
