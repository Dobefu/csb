package utils

import (
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/Dobefu/csb/cmd/database/query"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
	"github.com/Dobefu/csb/cmd/logger"
)

type AltLocale struct {
	Uid         string `json:"uid"`
	ContentType string `json:"content_type"`
	Locale      string `json:"locale"`
	Slug        string `json:"slug"`
	Url         string `json:"url"`
}

func GetAltLocales(entry structs.Route) ([]AltLocale, error) {
	rows, err := query.QueryRows(
		"routes",
		[]string{"uid", "contentType", "locale", "slug", "url"},
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

	var results []AltLocale

	for rows.Next() {
		var result AltLocale

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
