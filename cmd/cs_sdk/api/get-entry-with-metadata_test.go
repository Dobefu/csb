package api

import (
	"errors"
	"testing"
	"time"

	api_structs "github.com/Dobefu/csb/cmd/api/structs"
	"github.com/Dobefu/csb/cmd/api/utils"
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/stretchr/testify/assert"
)

func TestGetEntryWithMetadataSuccess(t *testing.T) {
	getEntry = func(route structs.Route) (map[string]interface{}, error) {
		return map[string]interface{}{"route_uid": route.Uid}, nil
	}

	getAltLocales = func(entry structs.Route, includeSitemapExcluded bool) ([]api_structs.AltLocale, error) {
		return []api_structs.AltLocale{{Uid: "en"}}, nil
	}

	getBreadcrumbs = func(entry structs.Route) ([]structs.Route, error) {
		return []structs.Route{{Uid: "uid"}}, nil
	}

	defer func() { getEntry = GetEntry }()
	defer func() { getAltLocales = utils.GetAltLocales }()
	defer func() { getBreadcrumbs = utils.GetBreadcrumbs }()

	entry, altLocales, breadcrumbs, err := GetEntryWithMetadata(structs.Route{Uid: "test_uid"})
	assert.Equal(t, map[string]interface{}{"route_uid": "test_uid"}, entry)
	assert.Equal(t, []api_structs.AltLocale([]api_structs.AltLocale{{Uid: "en", ContentType: "", Locale: "", Slug: "", Url: ""}}), altLocales)
	assert.Equal(t, []structs.Route([]structs.Route{{Id: "", Uid: "uid", Title: "", ContentType: "", Locale: "", Slug: "", Url: "", Parent: "", Version: 0, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), ExcludeSitemap: false, Published: false}}), breadcrumbs)
	assert.Equal(t, nil, err)
}

func TestGetEntryWithMetadataGetBreadcrumbsErr(t *testing.T) {
	getEntry = func(route structs.Route) (map[string]interface{}, error) {
		return map[string]interface{}{"route_uid": route.Uid}, nil
	}

	getAltLocales = func(entry structs.Route, includeSitemapExcluded bool) ([]api_structs.AltLocale, error) {
		return []api_structs.AltLocale{{Uid: "en"}}, nil
	}

	getBreadcrumbs = func(entry structs.Route) ([]structs.Route, error) {
		return nil, errors.New("")
	}

	defer func() { getEntry = GetEntry }()
	defer func() { getAltLocales = utils.GetAltLocales }()
	defer func() { getBreadcrumbs = utils.GetBreadcrumbs }()

	entry, altLocales, breadcrumbs, err := GetEntryWithMetadata(structs.Route{Uid: "test_uid"})
	assert.Equal(t, nil, entry)
	assert.Equal(t, []api_structs.AltLocale([]api_structs.AltLocale(nil)), altLocales)
	assert.Equal(t, []structs.Route([]structs.Route(nil)), breadcrumbs)
	assert.NotEqual(t, nil, err)
}

func TestGetEntryWithMetadataGetAltLocalesErr(t *testing.T) {
	getEntry = func(route structs.Route) (map[string]interface{}, error) {
		return map[string]interface{}{"route_uid": route.Uid}, nil
	}

	getAltLocales = func(entry structs.Route, includeSitemapExcluded bool) ([]api_structs.AltLocale, error) {
		return nil, errors.New("")
	}

	defer func() { getEntry = GetEntry }()
	defer func() { getAltLocales = utils.GetAltLocales }()

	entry, altLocales, breadcrumbs, err := GetEntryWithMetadata(structs.Route{Uid: "test_uid"})
	assert.Equal(t, nil, entry)
	assert.Equal(t, []api_structs.AltLocale([]api_structs.AltLocale(nil)), altLocales)
	assert.Equal(t, []structs.Route([]structs.Route(nil)), breadcrumbs)
	assert.NotEqual(t, nil, err)
}

func TestGetEntryWithMetadataGetEntryErr(t *testing.T) {
	getEntry = func(route structs.Route) (map[string]interface{}, error) {
		return nil, errors.New("")
	}

	defer func() { getEntry = GetEntry }()

	entry, altLocales, breadcrumbs, err := GetEntryWithMetadata(structs.Route{Uid: "test_uid"})
	assert.Equal(t, nil, entry)
	assert.Equal(t, []api_structs.AltLocale([]api_structs.AltLocale(nil)), altLocales)
	assert.Equal(t, []structs.Route([]structs.Route(nil)), breadcrumbs)
	assert.NotEqual(t, nil, err)
}
