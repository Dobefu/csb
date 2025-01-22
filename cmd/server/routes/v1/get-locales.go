package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dobefu/csb/cmd/api"
	"github.com/Dobefu/csb/cmd/server/utils"
)

var apiGetLocales = api.GetLocales
var utilsConstructOutput = utils.ConstructOutput
var utilsPrintError = utils.PrintError

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
