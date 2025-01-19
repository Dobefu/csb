package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dobefu/csb/cmd/api"
)

var apiGetTranslations = api.GetTranslations

func GetTranslations(w http.ResponseWriter, r *http.Request) {
	params, err := validationCheckRequiredQueryParams(
		r,
		"locale",
	)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utilsPrintError(w, err, false)
		return
	}

	locale := params["locale"].(string)

	translations, err := apiGetTranslations(locale)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utilsPrintError(w, err, true)
		return
	}

	output := utilsConstructOutput()
	output["data"] = translations

	json, err := json.Marshal(output)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utilsPrintError(w, err, true)
		return
	}

	fmt.Fprint(w, string(json))
}
