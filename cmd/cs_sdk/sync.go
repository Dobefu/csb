package cs_sdk

import (
	"errors"
	"fmt"
	"log"

	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/Dobefu/csb/cmd/database"
)

func Sync(reset bool) error {
	routes := make(map[string]structs.Route)

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

		err = addSyncRoutes(data, &routes)

		if err != nil {
			return err
		}

		err = processSyncData(routes)

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

func addSyncRoutes(data map[string]interface{}, routes *map[string]structs.Route) error {
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
		contentType := item["content_type_uid"].(string)
		parent := getParentUid(data)
		isPublished := hasPublishDetails
		id := fmt.Sprintf("%s%s", uid, locale)

		(*routes)[id] = structs.Route{
			Uid:         uid,
			ContentType: contentType,
			Locale:      locale,
			Slug:        slug,
			Url:         slug,
			Parent:      parent,
			Published:   isPublished,
		}
	}

	return nil
}

func getParentUid(data map[string]interface{}) string {
	parents, hasParents := data["parent"].([]interface{})

	if !hasParents || len(parents) <= 0 {
		return ""
	}

	parentData, hasParentData := parents[0].(map[string]interface{})

	if !hasParentData {
		return ""
	}

	parentUid, hasParentUid := parentData["uid"].(string)

	if !hasParentUid {
		return ""
	}

	return parentUid
}

func processSyncData(routes map[string]structs.Route) error {
	for idx, route := range routes {
		url := constructRouteUrl(route, routes)

		routes[idx] = structs.Route{
			Uid:         route.Uid,
			ContentType: route.ContentType,
			Locale:      route.Locale,
			Slug:        route.Slug,
			Url:         url,
			Parent:      route.Parent,
			Published:   route.Published,
		}

		err := database.SetRoute(routes[idx])

		if err != nil {
			return err
		}
	}

	return nil
}

func constructRouteUrl(route structs.Route, routes map[string]structs.Route) string {
	url := ""
	currentRoute := route
	depth := 0
	maxDepth := 10

	for {
		if depth > maxDepth {
			log.Printf("⚠️ Maximum nesting depth of %d exceeded in entry %s\n", maxDepth, route.Uid)
			return url
		}

		url = fmt.Sprintf("%s%s", currentRoute.Slug, url)
		parentUid := currentRoute.Parent

		if parentUid == "" {
			break
		}

		parentId := fmt.Sprintf("%s%s", parentUid, currentRoute.Locale)
		parent, hasParent := routes[parentId]

		if !hasParent {
			return url
		}

		if !parent.Published {
			log.Printf(
				"⚠️ The entry %s, a parent of %s, is unpublished. This will break the URL. Please be sure to publish it\n",
				parent.Uid,
				route.Uid,
			)
		}

		currentRoute = parent
		depth += 1
	}

	return url
}
