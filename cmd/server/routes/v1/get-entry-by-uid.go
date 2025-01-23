package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetEntryByUid(w http.ResponseWriter, r *http.Request) {
	params, err := validationCheckRequiredQueryParams(
		r,
		"uid",
		"locale",
	)

	if err != nil {
		utilsPrintError(w, err, false)
		return
	}

	uid := params["uid"].(string)
	locale := params["locale"].(string)

	entry, err := apiGetEntryByUid(uid, locale, false)

	if err != nil {
		utilsPrintError(w, err, true)
		return
	}

	csEntry, altLocales, breadcrumbs, err := csApiGetEntryWithMetadata(entry)

	if err != nil {
		utilsPrintError(w, err, true)
		return
	}

	output := utilsConstructEntryOutput(csEntry, altLocales, breadcrumbs)
	json, err := json.Marshal(output)

	if err != nil {
		utilsPrintError(w, err, true)
		return
	}

	fmt.Fprint(w, string(json))
}
