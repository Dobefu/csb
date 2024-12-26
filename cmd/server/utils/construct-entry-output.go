package utils

import (
	api_structs "github.com/Dobefu/csb/cmd/api/structs"
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
)

func ConstructEntryOutput(
	entry interface{},
	altLocales []api_structs.AltLocale,
	breadcrumbs []structs.Route,
) map[string]map[string]interface{} {
	output := ConstructOutput()

	output["data"]["entry"] = entry
	output["data"]["alt_locales"] = altLocales
	output["data"]["breadcrumbs"] = breadcrumbs

	return output
}
