package remote_sync

import (
	"errors"
	"testing"

	"github.com/Dobefu/csb/cmd/cs_sdk/functions"
	"github.com/stretchr/testify/assert"
)

func TestSyncSuccess(t *testing.T) {
	functionsSync = func(reset bool) error { return nil }
	defer func() { functionsSync = functions.Sync }()

	err := Sync(true)
	assert.NoError(t, err)
}

func TestSyncErr(t *testing.T) {
	functionsSync = func(reset bool) error { return errors.New("sync error") }
	defer func() { functionsSync = functions.Sync }()

	err := Sync(true)
	assert.EqualError(t, err, "sync error")
}
