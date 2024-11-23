package remote_sync

import "github.com/Dobefu/csb/cmd/cs_sdk"

func Sync(reset bool) error {
	return cs_sdk.Sync(reset)
}
