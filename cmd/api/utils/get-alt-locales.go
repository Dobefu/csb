package utils

import (
	api_structs "github.com/Dobefu/csb/cmd/api/structs"
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/Dobefu/csb/cmd/database/query"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
	"github.com/Dobefu/csb/cmd/logger"
)

func GetAltLocales(entry structs.Route) ([]api_structs.AltLocale, error) {
	rows, err := query.QueryRows(
		"routes",
		[]string{"uid", "content_type", "locale", "slug", "url"},
		[]db_structs.QueryWhere{
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
		},
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
