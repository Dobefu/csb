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

func GetEntryByUid(w http.ResponseWriter, r *http.Request) {
	params, err := validation.CheckRequiredQueryParams(
		r,
		"uid",
		"locale",
	)

	if err != nil {
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
		return
	}

	uid := params["uid"].(string)
	locale := params["locale"].(string)

	entry, err := api.GetEntryByUid(uid, locale, false)

	if err != nil {
		utils.PrintError(w, err, false)
		return
	}

	csEntry, altLocales, err := cs_api.GetEntryWithAltLocales(entry)

	if err != nil {
		utils.PrintError(w, err, false)
		return
	}

	output := utils.ConstructOutput()
	output["data"]["entry"] = csEntry
	output["data"]["alt_locales"] = altLocales

	json, err := json.Marshal(output)

	if err != nil {
		utils.PrintError(w, err, true)
		return
	}

	fmt.Fprint(w, string(json))
}
