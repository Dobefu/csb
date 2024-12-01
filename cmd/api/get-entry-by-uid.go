package api

import (
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
)

func GetEntryByUid(uid string, locale string, includeUnpublished bool) (structs.Route, error) {
	where := []db_structs.QueryWhere{
		{
			Name:  "uid",
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

	return GetEntryByFields(where)
}
