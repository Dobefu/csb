package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dobefu/csb/cmd/api"
)

var apiGetEntryByUrl = api.GetEntryByUrl

func GetEntryByUrl(w http.ResponseWriter, r *http.Request) {
	params, err := validationCheckRequiredQueryParams(
		r,
		"url",
		"locale",
	)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utilsPrintError(w, err, false)
		return
	}

	url := params["url"].(string)
	locale := params["locale"].(string)

	entry, err := apiGetEntryByUrl(url, locale, false)

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
