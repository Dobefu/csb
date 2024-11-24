package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSubCommands(t *testing.T) {
	os.Args[1] = "migrate:db"
	_, v := parseSubCommands()
	assert.NotEmpty(t, v.desc)

	os.Args[1] = "bogus"
	_, v = parseSubCommands()
	assert.Empty(t, v.desc)
}

func TestListSubCommands(t *testing.T) {
	listSubCommands()
}
