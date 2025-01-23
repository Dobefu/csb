package api

import (
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
)

func GetChildEntriesByUid(uid string, locale string, includeUnpublished bool) ([]structs.Route, error) {
	where := []db_structs.QueryWhere{
		{
			Name:  "parent",
			Value: uid,
		},
		{
			Name:  "locale",
			Value: locale,
		},
	}

	if !includeUnpublished {
		where = append(where, db_structs.QueryWhere{
			Name:  "published",
			Value: true,
		})
	}

	rows, err := queryQueryRows(
		"routes",
		[]string{"*"},
		where,
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
			&result.UpdatedAt,
			&result.ExcludeSitemap,
			&result.Published,
		)

		if err != nil {
			continue
		}

		results = append(results, result)
	}

	return results, nil
}
