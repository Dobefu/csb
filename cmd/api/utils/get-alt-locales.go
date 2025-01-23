package utils

import (
	api_structs "github.com/Dobefu/csb/cmd/api/structs"
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
	"github.com/Dobefu/csb/cmd/logger"
)

func GetAltLocales(entry structs.Route, includeSitemapExcluded bool) ([]api_structs.AltLocale, error) {
	where := []db_structs.QueryWhere{
		{
			Name:  "uid",
			Value: entry.Uid,
		},
		{
			Name:  "published",
			Value: true,
		},
		{
			Name:     "locale",
			Value:    entry.Locale,
			Operator: db_structs.NOT_EQUALS,
		},
	}

	if !includeSitemapExcluded {
		where = append(where, db_structs.QueryWhere{
			Name:     "exclude_sitemap",
			Value:    true,
			Operator: db_structs.NOT_EQUALS,
		})
	}

	rows, err := queryQueryRows(
		"routes",
		[]string{"uid", "content_type", "locale", "slug", "url"},
		where,
	)

	if err != nil {
		return nil, err
	}

	var results []api_structs.AltLocale

	for rows.Next() {
		var result api_structs.AltLocale

		err := rows.Scan(
			&result.Uid,
			&result.ContentType,
			&result.Locale,
			&result.Slug,
			&result.Url,
		)

		if err != nil {
			logger.Warning(err.Error())
			continue
		}

		results = append(results, result)
	}

	return results, nil
}
