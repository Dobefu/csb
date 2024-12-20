package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	setArgs("")
	main()

	setArgs("check:health", "--env-file=.env.test")
	main()

	setArgs("migrate:db", "--reset", "--verbose", "--env-file=.env.test")
	main()

	setArgs("remote:sync", "--reset", "--quiet", "--env-file=.env.test")
	main()

	setArgs("server", "--port=40000", "--env-file=.env.test")
	main()

	setArgs("create:content-type", "--name=\"Test content type\"", "--machine-name=\"test-content-type\"", "--dry-run", "--env-file=.env.test")
	main()
}

func setArgs(args ...string) {
	os.Args = []string{os.Args[0]}
	os.Args = append(os.Args, args...)
}

func TestListSubCommands(t *testing.T) {
	listSubCommands()
}
