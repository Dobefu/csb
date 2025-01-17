package utils

import (
	"github.com/Dobefu/csb/cmd/api"
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
)

var apiGetEntryByUid = api.GetEntryByUid
var apiGetEntryByUrl = api.GetEntryByUrl

func GetBreadcrumbs(entry structs.Route) ([]structs.Route, error) {
	results := []structs.Route{entry}
	currentEntry := entry

	var err error

	for currentEntry.Parent != "" {
		currentEntry, err = apiGetEntryByUid(currentEntry.Parent, entry.Locale, false)

		if err != nil {
			continue
		}

		results = append([]structs.Route{currentEntry}, results...)
	}

	if results[0].Url != "/" {
		homeEntry, err := apiGetEntryByUrl("/", entry.Locale, false)

		if err == nil {
			results = append([]structs.Route{homeEntry}, results...)
		}
	}

	return results, nil
}
