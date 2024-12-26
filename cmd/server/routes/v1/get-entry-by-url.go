package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dobefu/csb/cmd/api"
	cs_api "github.com/Dobefu/csb/cmd/cs_sdk/api"
	"github.com/Dobefu/csb/cmd/server/utils"
	"github.com/Dobefu/csb/cmd/server/validation"
)

func GetEntryByUrl(w http.ResponseWriter, r *http.Request) {
	params, err := validation.CheckRequiredQueryParams(
		r,
		"url",
		"locale",
	)

	if err != nil {
		utils.PrintError(w, err, false)
		return
	}

	url := params["url"].(string)
	locale := params["locale"].(string)

	entry, err := api.GetEntryByUrl(url, locale, false)

	if err != nil {
		utils.PrintError(w, err, false)
		return
	}

	csEntry, altLocales, breadcrumbs, err := cs_api.GetEntryWithMetadata(entry)

	if err != nil {
		utils.PrintError(w, err, false)
		return
	}

	output := utils.ConstructEntryOutput(csEntry, altLocales, breadcrumbs)
	json, err := json.Marshal(output)

	if err != nil {
		utils.PrintError(w, err, true)
		return
	}

	fmt.Fprint(w, string(json))
}
