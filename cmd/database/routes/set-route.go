package routes

import (
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/Dobefu/csb/cmd/cs_sdk/utils"
	"github.com/Dobefu/csb/cmd/database/query"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
)

var queryUpsert = query.Upsert
var utilsGenerateId = utils.GenerateId

func SetRoute(
	route structs.Route,
) error {
	err := queryUpsert("routes", getRouteValues(route))

	if err != nil {
		return err
	}

	return nil
}

func getRouteValues(route structs.Route) []db_structs.QueryValue {
	return []db_structs.QueryValue{
		{
			Name:  "id",
			Value: utilsGenerateId(route.Uid, route.Locale),
		},
		{
			Name:  "uid",
			Value: route.Uid,
		},
		{
			Name:  "title",
			Value: route.Title,
		},
		{
			Name:  "content_type",
			Value: route.ContentType,
		},
		{
			Name:  "locale",
			Value: route.Locale,
		},
		{
			Name:  "slug",
			Value: route.Slug,
		},
		{
			Name:  "url",
			Value: route.Url,
		},
		{
			Name:  "parent",
			Value: route.Parent,
		},
		{
			Name:  "version",
			Value: route.Version,
		},
		{
			Name:  "updated_at",
			Value: route.UpdatedAt,
		},
		{
			Name:  "exclude_sitemap",
			Value: route.ExcludeSitemap,
		},
		{
			Name:  "published",
			Value: route.Published,
		},
	}
}
