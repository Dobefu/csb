package api

import (
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/Dobefu/csb/cmd/database/query"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
)

func GetEntriesByFields(where []db_structs.QueryWhere) ([]structs.Route, error) {
	rows, err := query.QueryRows("routes", []string{"*"}, where)

	if err != nil {
		return nil, err
	}

	var results []structs.Route

	for rows.Next() {
		var result structs.Route

		err = rows.Scan(
			&result.Id,
			&result.Uid,
			&result.Title,
			&result.ContentType,
			&result.Locale,
			&result.Slug,
			&result.Url,
			&result.Parent,
			&result.ExcludeSitemap,
			&result.Published,
		)

		if err != nil {
			return results, err
		}
	}

	return results, nil
}
