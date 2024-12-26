package utils

import (
	"github.com/Dobefu/csb/cmd/api"
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
)

func GetBreadcrumbs(entry structs.Route) ([]structs.Route, error) {
	results := []structs.Route{entry}
	currentEntry := entry

	var err error

	for currentEntry.Parent != "" {
		currentEntry, err = api.GetEntryByUid(currentEntry.Parent, entry.Locale, false)

		if err != nil {
			continue
		}

		results = append([]structs.Route{currentEntry}, results...)
	}

	return results, nil
}
