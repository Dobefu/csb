package cs_sdk

import (
	"fmt"

	"github.com/Dobefu/csb/cmd/database"
)

func Sync(reset bool) error {
	paginationToken := ""

	for {
		var data map[string]interface{}

		syncToken := ""
		var err error

		if !reset {
			syncToken, err = database.GetState("sync_token")
		}

		if paginationToken != "" {
			path := fmt.Sprintf("stacks/sync?pagination_token=%s", paginationToken)
			data, err = Request(path, "GET")
		} else if err != nil || reset {
			path := fmt.Sprintf("stacks/sync?init=true")
			data, err = Request(path, "GET")
		} else {
			path := fmt.Sprintf("stacks/sync?sync_token=%s", syncToken)
			data, err = Request(path, "GET")
		}

		if err != nil {
			return err
		}

		newSyncToken, hasNewSyncToken := data["sync_token"].(string)

		if hasNewSyncToken {
			database.SetState("sync_token", newSyncToken)
		}

		processSyncData(data)

		var hasPaginationToken bool

		paginationToken, hasPaginationToken = data["pagination_token"].(string)

		if !hasPaginationToken {
			break
		}
	}

	return nil
}

func processSyncData(data map[string]interface{}) {

}
