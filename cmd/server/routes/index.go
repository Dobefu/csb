package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dobefu/csb/cmd/server/utils"
)

var utilsConstructOutput = utils.ConstructOutput

func Index(w http.ResponseWriter, r *http.Request, apiPath string) {
	// If not on the homepage, return a 404.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	output := utilsConstructOutput()
	output["data"]["api_endpoints"] = []string{apiPath}

	json, err := json.Marshal(output)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.PrintError(w, err, true)
		return
	}

	fmt.Fprint(w, string(json))
}
