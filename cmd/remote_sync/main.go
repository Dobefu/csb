package remote_sync

import "github.com/Dobefu/csb/cmd/cs_sdk"

func Sync() error {
	return cs_sdk.Sync()
}
