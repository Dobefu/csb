package cs_sdk

import (
	"errors"
	"fmt"

	"github.com/Dobefu/csb/cmd/database"
)

func Sync(reset bool) error {
	syncToken := ""
	paginationToken := ""

	for {
		data, err := getSyncData(paginationToken, reset, syncToken)

		if err != nil {
			return err
		}

		newSyncToken, hasNewSyncToken := data["sync_token"].(string)

		if hasNewSyncToken {
			database.SetState("sync_token", newSyncToken)
		}

		err = processSyncData(data)

		if err != nil {
			return err
		}

		var hasPaginationToken bool

		paginationToken, hasPaginationToken = data["pagination_token"].(string)

		if !hasPaginationToken {
			break
		}
	}

	return nil
}

func getSyncData(paginationToken string, reset bool, syncToken string) (map[string]interface{}, error) {
	var data map[string]interface{}
	var err error

	if !reset {
		syncToken, err = database.GetState("sync_token")
	}

	if paginationToken != "" {
		path := fmt.Sprintf("stacks/sync?pagination_token=%s", paginationToken)
		data, err = Request(path, "GET")
	} else if err != nil || reset {
		path := fmt.Sprintf("stacks/sync?init=true&type=entry_published,entry_unpublished")
		data, err = Request(path, "GET")
	} else {
		path := fmt.Sprintf("stacks/sync?sync_token=%s", syncToken)
		data, err = Request(path, "GET")
	}

	if err != nil {
		return nil, err
	}

	return data, nil
}

func processSyncData(data map[string]interface{}) error {
	items, hasItems := data["items"].([]interface{})

	if !hasItems {
		return errors.New("sync data has no items")
	}

	for _, item := range items {
		item := item.(map[string]interface{})
		data := item["data"].(map[string]interface{})

		publishDetails, hasPublishDetails := data["publish_details"].(map[string]interface{})

		if !hasPublishDetails {
			publishDetails = map[string]interface{}{
				"locale": data["locale"],
			}
		}

		slug, hasSlug := data["url"].(string)

		if !hasSlug {
			slug = ""
		}

		locale := publishDetails["locale"].(string)
		uid := data["uid"].(string)
		id := fmt.Sprintf("%s%s", uid, locale)
		contentType := item["content_type_uid"].(string)
		url := slug
		parent := ""
		isPublished := hasPublishDetails

		err := database.SetRoute(id, uid, contentType, locale, slug, url, parent, isPublished)

		if err != nil {
			return err
		}
	}

	return nil
}
