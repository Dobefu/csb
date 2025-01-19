package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dobefu/csb/cmd/api"
	cs_api "github.com/Dobefu/csb/cmd/cs_sdk/api"
	"github.com/Dobefu/csb/cmd/server/utils"
)

var apiGetEntryByUid = api.GetEntryByUid
var csApiGetEntryWithMetadata = cs_api.GetEntryWithMetadata
var utilsConstructEntryOutput = utils.ConstructEntryOutput

func GetEntryByUid(w http.ResponseWriter, r *http.Request) {
	params, err := validationCheckRequiredQueryParams(
		r,
		"uid",
		"locale",
	)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utilsPrintError(w, err, false)
		return
	}

	uid := params["uid"].(string)
	locale := params["locale"].(string)

	entry, err := apiGetEntryByUid(uid, locale, false)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utilsPrintError(w, err, true)
		return
	}

	csEntry, altLocales, breadcrumbs, err := csApiGetEntryWithMetadata(entry)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utilsPrintError(w, err, true)
		return
	}

	output := utilsConstructEntryOutput(csEntry, altLocales, breadcrumbs)
	json, err := json.Marshal(output)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utilsPrintError(w, err, true)
		return
	}

	fmt.Fprint(w, string(json))
}
