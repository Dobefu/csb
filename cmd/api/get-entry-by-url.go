package api

import (
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
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

	return GetEntryByFields(where)
}
