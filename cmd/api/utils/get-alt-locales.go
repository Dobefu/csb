package utils

import (
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/Dobefu/csb/cmd/database/query"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
)

func GetAltLocales(entry structs.Route) ([]structs.Route, error) {
	rows, err := query.QueryRows(
		"routes",
		[]string{"*"},
		[]db_structs.QueryWhere{
			{
				Name:  "uid",
				Value: entry.Uid,
			},
			{
				Name:  "published",
				Value: true,
			},
		},
	)

	if err != nil {
		return nil, err
	}

	var results []structs.Route

	for rows.Next() {
		var result structs.Route

		err := rows.Scan(
			&result.Id,
			&result.Uid,
			&result.ContentType,
			&result.Locale,
			&result.Slug,
			&result.Url,
			&result.Parent,
			&result.Published,
		)

		if err != nil {
			continue
		}

		if result.Locale == entry.Locale {
			continue
		}

		results = append(results, result)
	}

	return results, nil
}
