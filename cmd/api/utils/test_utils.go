package utils

import (
	"github.com/Dobefu/csb/cmd/api"
	"github.com/Dobefu/csb/cmd/database/query"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
)

var apiGetEntryByUid = api.GetEntryByUid
var apiGetEntryByUrl = api.GetEntryByUrl

type queryRows interface {
	Next() bool
	Scan(dest ...interface{}) error
}

var queryQueryRows = func(table string, columns []string, where []db_structs.QueryWhere) (queryRows, error) {
	return query.QueryRows(table, columns, where)
}
