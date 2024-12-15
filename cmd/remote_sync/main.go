package remote_sync

import (
	"github.com/Dobefu/csb/cmd/cs_sdk/functions"
)

func Sync(reset bool) error {
	return functions.Sync(reset)
}
