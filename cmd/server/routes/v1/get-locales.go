package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetLocales(w http.ResponseWriter, r *http.Request) {
	locales, err := apiGetLocales()

	if err != nil {
		utilsPrintError(w, err, false)
		return
	}

	output := utilsConstructOutput()
	output["data"] = locales

	json, err := json.Marshal(output)

	if err != nil {
		utilsPrintError(w, err, true)
		return
	}

	fmt.Fprint(w, string(json))
}
