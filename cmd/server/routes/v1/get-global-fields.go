package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dobefu/csb/cmd/api"
)

var apiGetGlobalFields = api.GetGlobalFields

func GetGlobalFields(w http.ResponseWriter, r *http.Request) {
	globalFields, err := apiGetGlobalFields()

	if err != nil {
		utilsPrintError(w, err, true)
		return
	}

	output := utilsConstructOutput()
	output["data"] = globalFields

	json, err := json.Marshal(output)

	if err != nil {
		utilsPrintError(w, err, true)
		return
	}

	fmt.Fprint(w, string(json))
}
