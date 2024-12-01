package api

import (
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/Dobefu/csb/cmd/database/query"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
)

func GetEntryByUrl(url string, locale string, includeUnpublished bool) (structs.Route, error) {
	where := []db_structs.QueryWhere{
		{
			Name:  "url",
			Value: url,
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

	row := query.QueryRow("routes", []string{"*"}, where)

	var result structs.Route

	err := row.Scan(
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
		return result, err
	}

	return result, nil
}
