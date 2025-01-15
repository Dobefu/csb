package functions

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Dobefu/csb/cmd/api"
	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/Dobefu/csb/cmd/cs_sdk/utils"
	"github.com/Dobefu/csb/cmd/database/assets"
	"github.com/Dobefu/csb/cmd/database/query"
	db_routes "github.com/Dobefu/csb/cmd/database/routes"
	"github.com/Dobefu/csb/cmd/database/state"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
	"github.com/Dobefu/csb/cmd/logger"
)

func Sync(reset bool) error {
	routes := make(map[string]structs.Route)

	if reset {
		logger.Info("Truncating the routes table")
		err := query.Truncate("routes")

		if err != nil {
			return err
		}
	}

	syncToken := ""
	paginationToken := ""

	for {
		data, err := getSyncData(paginationToken, reset, syncToken)

		if err != nil {
			return err
		}

		_, err = getNewSyncToken(data)

		if err != nil {
			return err
		}

		err = addAllRoutes(data, &routes)

		if err != nil {
			return err
		}

		err = addAllAssets(data)

		if err != nil {
			return err
		}

		var hasPaginationToken bool

		paginationToken, hasPaginationToken = data["pagination_token"].(string)

		if !hasPaginationToken {
			break
		}
	}

	logger.Info("Processing the sync items")
	err := processSyncData(routes)

	if err != nil {
		return err
	}

	logger.Info("Data sync completed successfully")

	return nil
}

func getNewSyncToken(data map[string]interface{}) (string, error) {
	newSyncToken, hasNewSyncToken := data["sync_token"].(string)

	if hasNewSyncToken {
		err := state.SetState("sync_token", newSyncToken)

		if err != nil {
			return "", err
		}
	}

	return newSyncToken, nil
}

func addAllAssets(data map[string]interface{}) error {
	items, hasItems := data["items"].([]interface{})

	if !hasItems {
		return errors.New("sync data has no items")
	}

	itemCount := len(items)

	for idx, item := range items {
		item := item.(map[string]interface{})

		if item["content_type_uid"].(string) != "sys_assets" {
			continue
		}

		logger.Info("Fetching item data (%d/%d)", (idx + 1), itemCount)

		assetData := item["data"].(map[string]interface{})
		publishDetails, hasPublishDetails := assetData["publish_details"].(map[string]interface{})

		if !hasPublishDetails {
			publishDetails = map[string]interface{}{
				"locale": assetData["locale"],
			}
		}

		parentUid, hasParentUid := assetData["parent_uid"].(string)

		if !hasParentUid {
			parentUid = ""
		}

		assetHeight, assetWidth := getAssetDimensions(assetData)
		filesize, err := strconv.Atoi(assetData["file_size"].(string))

		if err != nil {
			filesize = 0
		}

		err = assets.SetAsset(structs.Asset{
			Uid:         assetData["uid"].(string),
			Title:       assetData["title"].(string),
			ContentType: assetData["content_type"].(string),
			Locale:      publishDetails["locale"].(string),
			Url:         assetData["url"].(string),
			Parent:      parentUid,
			Filesize:    filesize,
			Height:      assetHeight,
			Width:       assetWidth,
			UpdatedAt:   getUpdatedAt(assetData),
			Published:   item["type"].(string) == "asset_published",
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func addAllRoutes(data map[string]interface{}, routes *map[string]structs.Route) error {
	err := addSyncRoutes(data, routes)

	if err != nil {
		return err
	}

	logger.Info("Adding child routes")
	err = addChildRoutes(routes)

	if err != nil {
		return err
	}

	logger.Info("Adding parent routes")
	err = addParentRoutes(routes)

	if err != nil {
		return err
	}

	return nil
}

func getSyncData(paginationToken string, reset bool, syncToken string) (map[string]interface{}, error) {
	var data map[string]interface{}
	var err error

	if !reset {
		syncToken, err = state.GetState("sync_token")
	}

	if paginationToken != "" {
		logger.Info("Getting a new sync page")
		path := fmt.Sprintf("stacks/sync?pagination_token=%s", paginationToken)
		data, err = cs_sdk.Request(path, "GET", nil, false)
	} else if err != nil || reset {
		logger.Info("Initialising a fresh sync")
		path := "stacks/sync?init=true&type=entry_published,entry_unpublished,entry_deleted,asset_published,asset_unpublished,asset_deleted"
		data, err = cs_sdk.Request(path, "GET", nil, false)
	} else {
		logger.Info("Syncing data using an existing sync token")
		path := fmt.Sprintf("stacks/sync?sync_token=%s", syncToken)
		data, err = cs_sdk.Request(path, "GET", nil, false)
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

	itemCount := len(items)

	for idx, item := range items {
		item := item.(map[string]interface{})

		if item["content_type_uid"].(string) == "sys_assets" {
			continue
		}

		logger.Info("Fetching item data (%d/%d)", (idx + 1), itemCount)

		data := item["data"].(map[string]interface{})
		publishDetails, hasPublishDetails := data["publish_details"].(map[string]interface{})

		if !hasPublishDetails {
			publishDetails = map[string]interface{}{
				"locale": data["locale"],
			}
		}

		uid := data["uid"].(string)
		title := getTitle(data)
		contentType := item["content_type_uid"].(string)
		locale := publishDetails["locale"].(string)
		slug := getSlug(data)
		parent := getParentUid(data)
		updatedAt := getUpdatedAt(data)
		excludeSitemap := getExcludeSitemap(data)
		isPublished := hasPublishDetails
		id := utils.GenerateId(uid, locale)

		logger.Verbose("Found entry: %s", uid)

		(*routes)[id] = structs.Route{
			Uid:            uid,
			Title:          title,
			ContentType:    contentType,
			Locale:         locale,
			Slug:           slug,
			Url:            slug,
			Parent:         parent,
			UpdatedAt:      updatedAt,
			ExcludeSitemap: excludeSitemap,
			Published:      isPublished,
		}
	}

	return nil
}

func getTitle(data map[string]interface{}) string {
	title, hasTitle := data["title"].(string)

	seo, hasSeo := data["seo"].(map[string]interface{})

	if !hasSeo {
		if hasTitle {
			return title
		}

		return ""
	}

	seoTitle, hasSeoTitle := seo["title"]

	if !hasSeoTitle || seoTitle == "" {
		if hasTitle {
			return title
		}

		return ""
	}

	return seoTitle.(string)
}

func getSlug(data map[string]interface{}) string {
	slug, hasSlug := data["url"].(string)

	if !hasSlug {
		slug = ""
	}

	return slug
}

func addChildRoutes(routes *map[string]structs.Route) error {
	for _, route := range *routes {
		err := addRouteChildren(route, routes, 0)

		if err != nil {
			continue
		}
	}

	return nil
}

func addParentRoutes(routes *map[string]structs.Route) error {
	for _, route := range *routes {
		err := addRouteParents(route, routes, 0)

		if err != nil {
			continue
		}
	}

	return nil
}

func addRouteChildren(route structs.Route, routes *map[string]structs.Route, depth uint8) error {
	var maxDepth uint8 = 10

	if depth > maxDepth {
		return errors.New("potential infinite loop detected")
	}

	childRoutes, err := api.GetChildEntriesByUid(route.Uid, route.Locale, true)

	if err != nil {
		return err
	}

	for _, childRoute := range childRoutes {
		if childRoute.Uid == "" {
			continue
		}

		id := utils.GenerateId(childRoute.Uid, childRoute.Locale)
		(*routes)[id] = childRoute

		err = addRouteChildren(childRoute, routes, depth+1)

		if err != nil {
			logger.Warning("Error getting a child route for %s: %s", id, err.Error())
			return err
		}
	}

	return nil
}

func addRouteParents(route structs.Route, routes *map[string]structs.Route, depth uint8) error {
	var maxDepth uint8 = 10

	if depth > maxDepth {
		return errors.New("potential infinite loop detected")
	}

	parentId := utils.GenerateId(route.Parent, route.Locale)
	parentRoute := (*routes)[parentId]

	var err error

	// If the parent page cannot be found in the routes, check the database.
	if parentRoute.Uid == "" {
		parentRoute, err = api.GetEntryByUid(route.Parent, route.Locale, true)

		// If there is no parent, there will be an error.
		// This is expected, since the database will not have any results.
		// When this happens, we can just return without any error.
		if err != nil {
			return nil
		}
	}

	id := utils.GenerateId(parentRoute.Uid, parentRoute.Locale)
	(*routes)[id] = parentRoute

	err = addRouteParents(parentRoute, routes, depth+1)

	if err != nil {
		logger.Warning("Error getting a parent route for %s: %s", id, err.Error())
		return err
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

func getAssetDimensions(asset map[string]interface{}) (int, int) {
	dimension, hasDimension := asset["dimension"].(map[string]interface{})

	if !hasDimension {
		return 0, 0
	}

	return int(dimension["height"].(float64)), int(dimension["width"].(float64))
}

func getUpdatedAt(data map[string]interface{}) time.Time {
	updatedAt, hasUpdatedAt := data["updated_at"]

	if !hasUpdatedAt {
		return time.Now()
	}

	datetime, err := time.Parse(time.RFC3339, updatedAt.(string))

	if err != nil {
		return time.Now()
	}

	return datetime
}

func getExcludeSitemap(data map[string]interface{}) bool {
	seo, hasSeo := data["seo"].(map[string]interface{})

	if !hasSeo {
		return false
	}

	excludeSitemap, hasExcludeSitemap := seo["exclude_sitemap"]

	if !hasExcludeSitemap {
		return false
	}

	return excludeSitemap.(bool)
}

func processSyncData(routes map[string]structs.Route) error {
	routeCount := len(routes)
	idx := 0

	for uid, route := range routes {
		logger.Info("Processing entry %s (%d/%d)", route.Uid, (idx + 1), routeCount)
		url := constructRouteUrl(route, routes)

		if url == "" {
			if route.ContentType == "translations" {
				processTranslations(route)
			}

			idx += 1
			continue
		}

		routes[uid] = structs.Route{
			Uid:            route.Uid,
			Title:          route.Title,
			ContentType:    route.ContentType,
			Locale:         route.Locale,
			Slug:           route.Slug,
			Url:            url,
			Parent:         route.Parent,
			UpdatedAt:      route.UpdatedAt,
			ExcludeSitemap: route.ExcludeSitemap,
			Published:      route.Published,
		}

		err := db_routes.SetRoute(routes[uid])

		if err != nil {
			return err
		}

		idx += 1
	}

	return nil
}

func processTranslations(route structs.Route) {
	path := fmt.Sprintf(
		"content_types/%s/entries/%s?locale=%s",
		route.ContentType,
		route.Uid,
		route.Locale,
	)

	resp, err := cs_sdk.Request(path, "GET", nil, false)

	if err != nil {
		return
	}

	entry, hasEntry := resp["entry"].(map[string]interface{})

	if !hasEntry {
		return
	}

	translations, hasTranslations := entry["translations"]

	if !hasTranslations {
		return
	}

	category, hasCategory := entry["category"]

	if !hasCategory {
		return
	}

	for _, translation := range translations.([]interface{}) {
		source := translation.(map[string]interface{})["source"]
		err = query.Upsert("translations", []db_structs.QueryValue{
			{
				Name:  "id",
				Value: fmt.Sprintf("%s%s", utils.GenerateId(route.Uid, route.Locale), source),
			},
			{
				Name:  "uid",
				Value: route.Uid,
			},
			{
				Name:  "source",
				Value: source,
			},
			{
				Name:  "translation",
				Value: translation.(map[string]interface{})["translation"],
			},
			{
				Name:  "category",
				Value: category,
			},
			{
				Name:  "locale",
				Value: route.Locale,
			},
		})

		if err != nil {
			logger.Warning(err.Error())
		}
	}
}

func constructRouteUrl(route structs.Route, routes map[string]structs.Route) string {
	url := ""

	currentRoute := route

	var depth uint8 = 0
	var maxDepth uint8 = 10

	for {
		if depth > maxDepth {
			logger.Warning("Maximum nesting depth of %d exceeded in entry %s", maxDepth, route.Uid)
			return url
		}

		url = fmt.Sprintf("%s%s", currentRoute.Slug, url)
		parentUid := currentRoute.Parent

		if parentUid == "" {
			break
		}

		parentId := utils.GenerateId(parentUid, currentRoute.Locale)
		parent, hasParent := routes[parentId]

		if !hasParent {
			return url
		}

		if !parent.Published {
			logger.Warning(
				"The entry %s, a parent of %s, is unpublished. This will break the URL. Please be sure to publish it",
				parent.Uid,
				route.Uid,
			)
		}

		currentRoute = parent
		depth += 1
	}

	return strings.ReplaceAll(url, "//", "/")
}
