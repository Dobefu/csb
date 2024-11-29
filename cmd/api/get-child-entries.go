package api

import (
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/Dobefu/csb/cmd/database/query"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
)

func GetChildEntriesByUid(uid string, locale string) ([]structs.Route, error) {
	rows, err := query.QueryRows(
		"routes",
		[]string{"*"},
		[]db_structs.QueryWhere{
			{
				Name:  "parent",
				Value: uid,
			},
			{
				Name:  "locale",
				Value: locale,
			},
		},
	)

	if err != nil {
		return []structs.Route{}, err
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

		results = append(results, result)
	}

	return results, nil
}
