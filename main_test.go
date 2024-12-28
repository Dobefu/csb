package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	const env = "--env-file=.env.test"

	setArgs()
	main()

	setArgs("check:health", env)
	main()

	setArgs("migrate:db", "--reset", "--verbose", env)
	main()

	setArgs("remote:setup", "--quiet", env)
	main()

	setArgs("remote:sync", "--reset", "--quiet", env)
	main()

	setArgs("server", "--port=40000", env)
	main()

	setArgs("create:content-type", "--name=\"Test content type\"", "--machine-name=\"test-content-type\"", "--dry-run", env)
	main()
}

func setArgs(args ...string) {
	os.Args = []string{os.Args[0]}

	if len(args) > 0 {
		os.Args = append(os.Args, args...)
	}
}

func TestListSubCommands(t *testing.T) {
	listSubCommands()
}
