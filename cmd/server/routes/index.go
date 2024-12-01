package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dobefu/csb/cmd/logger"
	"github.com/Dobefu/csb/cmd/server/utils"
)

func Index(w http.ResponseWriter, r *http.Request, apiPath string) {
	// If not on the homepage, return a 404.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	output := utils.ConstructOutput()
	output["data"]["api_endpoints"] = []string{apiPath}

	json, err := json.Marshal(output)

	if err != nil {
		logger.Error(err.Error())
		utils.PrintError(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}
