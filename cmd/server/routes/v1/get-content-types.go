package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetContentTypes(w http.ResponseWriter, r *http.Request) {
	contentTypes, err := apiGetContentTypes()

	if err != nil {
		utilsPrintError(w, err, true)
		return
	}

	output := utilsConstructOutput()
	output["data"] = contentTypes

	json, err := json.Marshal(output)

	if err != nil {
		utilsPrintError(w, err, true)
		return
	}

	fmt.Fprint(w, string(json))
}
