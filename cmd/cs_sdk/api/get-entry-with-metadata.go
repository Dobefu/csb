package api

import (
	api_structs "github.com/Dobefu/csb/cmd/api/structs"
	"github.com/Dobefu/csb/cmd/api/utils"
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
)

var getEntry = GetEntry
var getAltLocales = utils.GetAltLocales
var getBreadcrumbs = utils.GetBreadcrumbs

func GetEntryWithMetadata(route structs.Route) (interface{}, []api_structs.AltLocale, []structs.Route, error) {
	entry, err := getEntry(route)

	if err != nil {
		return nil, nil, nil, err
	}

	altLocales, err := getAltLocales(route, true)

	if err != nil {
		return nil, nil, nil, err
	}

	breadcrumbs, err := getBreadcrumbs(route)

	if err != nil {
		return nil, nil, nil, err
	}

	return entry, altLocales, breadcrumbs, nil
}
