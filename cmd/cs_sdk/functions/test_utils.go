package functions

import (
	"github.com/Dobefu/csb/cmd/api"
	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/Dobefu/csb/cmd/cs_sdk/utils"
	"github.com/Dobefu/csb/cmd/database/assets"
	"github.com/Dobefu/csb/cmd/database/query"
	db_routes "github.com/Dobefu/csb/cmd/database/routes"
	"github.com/Dobefu/csb/cmd/database/state"
)

var queryTruncate = query.Truncate
var queryUpsert = query.Upsert
var stateSetState = state.SetState
var stateGetState = state.GetState
var csSdkRequest = cs_sdk.Request
var utilsGenerateId = utils.GenerateId
var apiGetChildEntriesByUid = api.GetChildEntriesByUid
var apiGetEntryByUid = api.GetEntryByUid
var dbRoutesSetRoute = db_routes.SetRoute
var assetsSetAsset = assets.SetAsset
