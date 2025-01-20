package functions

import (
	"errors"
	"testing"

	"github.com/Dobefu/csb/cmd/api"
	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/Dobefu/csb/cmd/cs_sdk/utils"
	"github.com/Dobefu/csb/cmd/database/assets"
	"github.com/Dobefu/csb/cmd/database/query"
	db_routes "github.com/Dobefu/csb/cmd/database/routes"
	"github.com/Dobefu/csb/cmd/database/state"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
	"github.com/stretchr/testify/assert"
)

func setupTestSync() func() {
	queryTruncate = func(table string) error { return nil }
	queryUpsert = func(table string, values []db_structs.QueryValue) error { return nil }
	stateSetState = func(name string, value string) error { return nil }
	stateGetState = func(name string) (string, error) { return "", nil }

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		data := map[string]interface{}{
			"items": []interface{}{
				map[string]interface{}{
					"content_type_uid": "page",
					"data": map[string]interface{}{
						"uid":    "test-uid-route",
						"url":    "/",
						"locale": "en",
					},
				},
				map[string]interface{}{
					"content_type_uid": "sys_assets",
					"type":             "asset_published",
					"data": map[string]interface{}{
						"uid":    "test-uid-asset",
						"locale": "en",
					},
				},
			},
			"sync_token": "sync-token",
		}

		return data, nil
	}

	utilsGenerateId = func(uid string, locale string) string { return "" }

	apiGetChildEntriesByUid = func(uid string, locale string, includeUnpublished bool) ([]structs.Route, error) {
		return []structs.Route{}, nil
	}

	apiGetEntryByUid = func(uid string, locale string, includeUnpublished bool) (structs.Route, error) {
		return structs.Route{}, nil
	}

	dbRoutesSetRoute = func(route structs.Route) error { return nil }
	assetsSetAsset = func(asset structs.Asset) error { return nil }

	cleanup := func() {
		queryTruncate = query.Truncate
		queryUpsert = query.Upsert
		stateSetState = state.SetState
		stateGetState = state.GetState
		csSdkRequest = cs_sdk.Request
		utilsGenerateId = utils.GenerateId
		apiGetChildEntriesByUid = api.GetChildEntriesByUid
		apiGetEntryByUid = api.GetEntryByUid
		dbRoutesSetRoute = db_routes.SetRoute
		assetsSetAsset = assets.SetAsset
	}

	return cleanup
}

func TestSyncResetSuccess(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	err := Sync(true)
	assert.NoError(t, err)
}

func TestSyncResetErrTruncate(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	queryTruncate = func(table string) error {
		return errors.New("truncate failed")
	}

	err := Sync(true)
	assert.EqualError(t, err, "truncate failed")
}

func TestSyncNoResetSuccess(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	err := Sync(false)
	assert.NoError(t, err)
}

func TestSyncNoResetErrCsSdkRequest(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		return nil, errors.New("SDK request failed")
	}

	err := Sync(false)
	assert.EqualError(t, err, "SDK request failed")
}

func TestSyncNoResetErrSetState(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	stateSetState = func(name string, value string) error {
		return errors.New("error setting state")
	}

	err := Sync(false)
	assert.EqualError(t, err, "error setting state")
}

func TestSyncNoResetErrAddAllRoutes(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		return map[string]interface{}{}, nil
	}

	err := Sync(false)
	assert.EqualError(t, err, "sync data has no items")
}

func TestSyncNoResetErrAddAssets(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	assetsSetAsset = func(asset structs.Asset) error {
		return errors.New("failed setting asset")
	}

	err := Sync(false)
	assert.EqualError(t, err, "failed setting asset")
}

func TestSyncNoResetErrProcessSyncData(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	dbRoutesSetRoute = func(route structs.Route) error {
		return errors.New("failed setting route")
	}

	err := Sync(false)
	assert.EqualError(t, err, "failed setting route")
}
