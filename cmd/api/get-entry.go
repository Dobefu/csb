package api

import (
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/Dobefu/csb/cmd/database/query"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
)

func GetEntry(uid string, locale string) (structs.Route, error) {
	row := query.QueryRow("routes", []string{"*"}, []db_structs.QueryWhere{
		{
			Name:  "uid",
			Value: uid,
		},
		{
			Name:  "locale",
			Value: locale,
		},
	})

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
