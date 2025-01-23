package v1

import (
	"github.com/Dobefu/csb/cmd/api"
	api_utils "github.com/Dobefu/csb/cmd/api/utils"
	cs_api "github.com/Dobefu/csb/cmd/cs_sdk/api"
	"github.com/Dobefu/csb/cmd/cs_sdk/functions"
	"github.com/Dobefu/csb/cmd/database/query"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
	"github.com/Dobefu/csb/cmd/server/utils"
	"github.com/Dobefu/csb/cmd/server/validation"
)

var validationCheckRequiredQueryParams = validation.CheckRequiredQueryParams
var apiGetContentType = api.GetContentType
var apiGetContentTypes = api.GetContentTypes
var apiGetEntryByUid = api.GetEntryByUid
var apiGetEntryByUrl = api.GetEntryByUrl
var csApiGetEntryWithMetadata = cs_api.GetEntryWithMetadata
var utilsConstructEntryOutput = utils.ConstructEntryOutput
var apiGetGlobalFields = api.GetGlobalFields
var apiGetLocales = api.GetLocales
var utilsConstructOutput = utils.ConstructOutput
var utilsPrintError = utils.PrintError
var apiUtilsGetAltLocales = api_utils.GetAltLocales
var apiGetTranslations = api.GetTranslations
var functionsSync = functions.Sync

var queryQueryRows = func(table string, columns []string, where []db_structs.QueryWhere) (queryRows, error) {
	return query.QueryRows(table, columns, where)
}

type queryRows interface {
	Next() bool
	Scan(dest ...interface{}) error
}
