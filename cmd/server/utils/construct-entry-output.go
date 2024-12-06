package utils

import (
	"github.com/Dobefu/csb/cmd/api/structs"
)

func ConstructEntryOutput(entry interface{}, altLocales []structs.AltLocale) map[string]map[string]interface{} {
	output := ConstructOutput()

	output["data"]["entry"] = entry
	output["data"]["alt_locales"] = altLocales

	return output
}
