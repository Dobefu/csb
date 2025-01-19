package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dobefu/csb/cmd/api"
)

var apiGetContentTypes = api.GetContentTypes

func GetContentTypes(w http.ResponseWriter, r *http.Request) {
	contentTypes, err := apiGetContentTypes()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utilsPrintError(w, err, false)
		return
	}

	output := utilsConstructOutput()
	output["data"] = contentTypes

	json, err := json.Marshal(output)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utilsPrintError(w, err, true)
		return
	}

	fmt.Fprint(w, string(json))
}
