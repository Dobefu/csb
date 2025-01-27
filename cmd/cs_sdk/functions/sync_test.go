package functions

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/Dobefu/csb/cmd/api"
	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/Dobefu/csb/cmd/cs_sdk/utils"
	"github.com/Dobefu/csb/cmd/database/assets"
	"github.com/Dobefu/csb/cmd/database/query"
	db_routes "github.com/Dobefu/csb/cmd/database/routes"
	"github.com/Dobefu/csb/cmd/database/state"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
	"github.com/Dobefu/csb/cmd/logger"
	"github.com/stretchr/testify/assert"
)

func setupTestSync() func() {
	queryTruncate = func(table string) error { return nil }
	queryUpsert = func(table string, values []db_structs.QueryValue) error { return nil }
	stateSetState = func(name string, value string) error { return nil }
	stateGetState = func(name string) (string, error) { return "", nil }
	loggerInfo = func(format string, a ...any) string { return "" }
	loggerWarning = func(format string, a ...any) string { return "" }
	loggerVerbose = func(format string, a ...any) string { return "" }

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

	utilsGenerateId = func(uid string, locale string) string {
		return fmt.Sprintf("%s%s", uid, locale)
	}

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
		loggerInfo = logger.Info
		loggerWarning = logger.Warning
		loggerVerbose = logger.Verbose
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

func TestAddAllAssetsErrNoItems(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	err := addAllAssets(make(map[string]interface{}))
	assert.EqualError(t, err, "sync data has no items")
}

func TestAddAllRoutesErrAddRouteChildren(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	apiGetChildEntriesByUid = func(uid string, locale string, includeUnpublished bool) ([]structs.Route, error) {
		return nil, errors.New("cannot get child entries")
	}

	err := addAllRoutes(
		map[string]interface{}{
			"items": []interface{}{
				map[string]interface{}{
					"content_type_uid": "page",
					"data": map[string]interface{}{
						"uid":    "test-uid-route",
						"locale": "en",
					},
				},
			},
		},
		&(map[string]structs.Route{}),
	)

	assert.EqualError(t, err, "cannot get child entries")
}

func TestAddAllRoutesErrAddRouteParents(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	entries := map[string]interface{}{"items": make([]interface{}, 11)}

	for i := 0; i < 11; i++ {
		entries["items"].([]interface{})[i] = map[string]interface{}{
			"content_type_uid": "page",
			"data": map[string]interface{}{
				"uid":    fmt.Sprintf("uid-%d", i),
				"locale": "en",
				"parent": []interface{}{map[string]interface{}{"uid": fmt.Sprintf("uid-%d", i-1)}},
			},
		}
	}

	err := addAllRoutes(
		entries,
		&(map[string]structs.Route{}),
	)

	assert.EqualError(t, err, "potential infinite loop detected")
}

func TestGetSyncDataPaginationToken(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	data, err := getSyncData("test-token", false, "")

	assert.NoError(t, err)
	assert.NotNil(t, data)
}

func TestGetFilesizeSuccess(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	filesize := getFilesize(map[string]interface{}{"file_size": "100"})

	assert.Equal(t, 100, filesize)
}

func TestGetFilesizeErrAtoi(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	filesize := getFilesize(map[string]interface{}{"file_size": "bogus"})

	assert.Equal(t, 0, filesize)
}

func TestGetTitleSuccessTitle(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	title := getTitle(map[string]interface{}{
		"title": "test-title",
	})

	assert.Equal(t, "test-title", title)
}

func TestGetTitleSuccessSeoTitle(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	title := getTitle(map[string]interface{}{
		"title": "test-title",
		"seo": map[string]interface{}{
			"title": "test-seo-title",
		},
	})

	assert.Equal(t, "test-seo-title", title)
}

func TestGetTitleErrSeoEmpty(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	title := getTitle(map[string]interface{}{
		"title": "test-title",
		"seo":   map[string]interface{}{},
	})

	assert.Equal(t, "test-title", title)
}

func TestGetTitleErrSeoEmptyNoTitle(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	title := getTitle(map[string]interface{}{
		"seo": map[string]interface{}{},
	})

	assert.Equal(t, "", title)
}

func TestAddRouteChildrenErrMaxDepth(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	route := structs.Route{
		Uid:    "test-uid",
		Locale: "en",
	}

	apiGetChildEntriesByUid = func(uid string, locale string, includePages bool) ([]structs.Route, error) {
		return []structs.Route{route}, nil
	}

	err := addRouteChildren(
		structs.Route{},
		&(map[string]structs.Route{}),
		0,
	)

	assert.EqualError(t, err, "potential infinite loop detected")
}

func TestAddRouteChildrenErrNoUid(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	apiGetChildEntriesByUid = func(uid string, locale string, includePages bool) ([]structs.Route, error) {
		return []structs.Route{{}}, nil
	}

	err := addRouteChildren(
		structs.Route{},
		&(map[string]structs.Route{}),
		0,
	)

	assert.NoError(t, err)
}

func TestAddRouteParentsNoParentUid(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	apiGetEntryByUid = func(uid string, locale string, includeUnpublished bool) (structs.Route, error) {
		return structs.Route{}, errors.New("cannot get entry by UID")
	}

	err := addRouteParents(
		structs.Route{Parent: "test-parent-uid"},
		&(map[string]structs.Route{}),
		0,
	)

	assert.NoError(t, err)
}

func TestGetParentUidNoParentData(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	apiGetEntryByUid = func(uid string, locale string, includeUnpublished bool) (structs.Route, error) {
		return structs.Route{}, errors.New("cannot get entry by UID")
	}

	parentUid := getParentUid(
		map[string]interface{}{
			"parent": []interface{}{[]interface{}{}},
		},
	)

	assert.Equal(t, "", parentUid)
}

func TestGetParentUidNoParentUid(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	apiGetEntryByUid = func(uid string, locale string, includeUnpublished bool) (structs.Route, error) {
		return structs.Route{}, errors.New("cannot get entry by UID")
	}

	parentUid := getParentUid(
		map[string]interface{}{
			"parent": []interface{}{map[string]interface{}{}},
		},
	)

	assert.Equal(t, "", parentUid)
}

func TestGetVersionSuccess(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	version := getVersion(map[string]interface{}{"_version": 1.0})

	assert.Equal(t, 1, version)
}

func TestGetAssetDimensionsSuccess(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	height, width := getAssetDimensions(map[string]interface{}{
		"dimension": map[string]interface{}{
			"height": 100.0,
			"width":  200.0,
		},
	})

	assert.Equal(t, 100, height)
	assert.Equal(t, 200, width)
}

func TestGetUpdatedAtSuccess(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	updatedAt := getUpdatedAt(map[string]interface{}{"updated_at": "2000-01-01T00:00:00+00:00"})

	assert.Equal(t, int64(946684800), updatedAt.Unix())
}

func TestGetUpdatedAtErrInvalid(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	updatedAt := getUpdatedAt(map[string]interface{}{"updated_at": "bogus"})

	assert.Equal(t, time.Now().Unix(), updatedAt.Unix())
}

func TestGetExcludeSitemapSuccess(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	excludeSitemap := getExcludeSitemap(map[string]interface{}{
		"seo": map[string]interface{}{
			"exclude_sitemap": true,
		},
	})

	assert.Equal(t, true, excludeSitemap)
}

func TestGetExcludeSitemapEmptySeo(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	excludeSitemap := getExcludeSitemap(map[string]interface{}{
		"seo": map[string]interface{}{},
	})

	assert.Equal(t, false, excludeSitemap)
}

func TestProcessSyncDataTranslationsSuccess(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		data := map[string]interface{}{
			"entry": map[string]interface{}{
				"content_type_uid": "translations",
				"category":         "category",
				"data": map[string]interface{}{
					"uid":    "translation-0",
					"locale": "en",
				},
				"translations": []interface{}{map[string]interface{}{
					"source":      "source",
					"translation": "translation",
				}},
			},
		}

		return data, nil
	}

	err := processSyncData(map[string]structs.Route{
		"t0": {
			Uid:         "translation-0",
			ContentType: "translations",
		},
	})

	assert.NoError(t, err)
}

func TestProcessTranslationsCsSdkGet(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		return nil, errors.New("cannot get entries")
	}

	processTranslations(structs.Route{
		Uid:         "translation-0",
		ContentType: "translations",
	})

	assert.Nil(t, nil)
}

func TestProcessTranslationsNoEntry(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		return map[string]interface{}{}, nil
	}

	processTranslations(structs.Route{
		Uid:         "translation-0",
		ContentType: "translations",
	})

	assert.Nil(t, nil)
}

func TestProcessTranslationsNoTranslations(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		data := map[string]interface{}{
			"entry": map[string]interface{}{
				"content_type_uid": "translations",
				"category":         "category",
				"data": map[string]interface{}{
					"uid":    "translation-invalid",
					"locale": "en",
				},
			},
		}

		return data, nil
	}

	processTranslations(structs.Route{
		Uid:         "translation-0",
		ContentType: "translations",
	})

	assert.Nil(t, nil)
}

func TestProcessTranslationsNoCategory(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		data := map[string]interface{}{
			"entry": map[string]interface{}{
				"content_type_uid": "translations",
				"data": map[string]interface{}{
					"uid":    "translation-invalid",
					"locale": "en",
				},
				"translations": []interface{}{map[string]interface{}{
					"source":      "source",
					"translation": "translation",
				}},
			},
		}

		return data, nil
	}

	processTranslations(structs.Route{
		Uid:         "translation-0",
		ContentType: "translations",
	})

	assert.Nil(t, nil)
}

func TestProcessSyncDataTranslationsErrUpsert(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		data := map[string]interface{}{
			"entry": map[string]interface{}{
				"content_type_uid": "translations",
				"category":         "category",
				"data": map[string]interface{}{
					"uid":    "translation-0",
					"locale": "en",
				},
				"translations": []interface{}{map[string]interface{}{
					"source":      "source",
					"translation": "translation",
				}},
			},
		}

		return data, nil
	}

	queryUpsert = func(table string, values []db_structs.QueryValue) error {
		return errors.New("cannot upsert translation")
	}

	err := processSyncData(map[string]structs.Route{
		"t0": {
			Uid:         "translation-0",
			ContentType: "translations",
		},
	})

	assert.NoError(t, err)
}

func TestConstructRouteUrlErrMaxDepth(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	url := constructRouteUrl(
		structs.Route{
			Uid:    "r0en",
			Locale: "en",
			Parent: "p0",
		},
		map[string]structs.Route{
			"r0en": {
				Uid:    "r0en",
				Locale: "en",
				Parent: "p0",
			},
			"p0en": {
				Uid:    "p0en",
				Locale: "en",
				Parent: "r0",
			},
		},
	)

	assert.Equal(t, "", url)
}

func TestConstructRouteUrlNoParent(t *testing.T) {
	cleanup := setupTestSync()
	defer cleanup()

	url := constructRouteUrl(
		structs.Route{
			Uid:    "r0",
			Locale: "en",
			Parent: "p0",
		},
		map[string]structs.Route{
			"r0en": {
				Uid:    "r0en",
				Locale: "en",
				Parent: "p0",
			},
		},
	)

	assert.Equal(t, "", url)
}
