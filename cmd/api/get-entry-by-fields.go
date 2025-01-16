package api

import (
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/Dobefu/csb/cmd/database/query"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
)

func GetEntryByFields(where []db_structs.QueryWhere) (structs.Route, error) {
	row := query.QueryRow("routes", []string{"*"}, where)

	var result structs.Route

	err := row.Scan(
		&result.Id,
		&result.Uid,
		&result.Title,
		&result.ContentType,
		&result.Locale,
		&result.Slug,
		&result.Url,
		&result.Parent,
		&result.Version,
		&result.UpdatedAt,
		&result.ExcludeSitemap,
		&result.Published,
	)

	if err != nil {
		return result, err
	}

	return result, nil
}
