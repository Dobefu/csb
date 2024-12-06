package utils

import (
	api_structs "github.com/Dobefu/csb/cmd/api/structs"
	"github.com/Dobefu/csb/cmd/api/utils"
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
)

func GetEntryWithAltLocales(route structs.Route) (interface{}, []api_structs.AltLocale, error) {
	entry, err := GetEntry(route)

	if err != nil {
		return nil, nil, err
	}

	altLocales, err := utils.GetAltLocales(route)

	if err != nil {
		return nil, nil, err
	}

	return entry, altLocales, nil
}
