package remote_sync

import (
	"github.com/Dobefu/csb/cmd/cs_sdk/functions"
)

var functionsSync = functions.Sync

func Sync(reset bool) error {
	err := functionsSync(reset)

	if err != nil {
		return err
	}

	return nil
}
