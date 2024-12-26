package api

import (
	api_structs "github.com/Dobefu/csb/cmd/api/structs"
	"github.com/Dobefu/csb/cmd/api/utils"
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
)

func GetEntryWithMetadata(route structs.Route) (interface{}, []api_structs.AltLocale, []structs.Route, error) {
	entry, err := GetEntry(route)

	if err != nil {
		return nil, nil, nil, err
	}

	altLocales, err := utils.GetAltLocales(route, true)

	if err != nil {
		return nil, nil, nil, err
	}

	breadcrumbs, err := utils.GetBreadcrumbs(route)

	if err != nil {
		return nil, nil, nil, err
	}

	return entry, altLocales, breadcrumbs, nil
}
