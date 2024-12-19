package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dobefu/csb/cmd/api"
	"github.com/Dobefu/csb/cmd/server/utils"
)

func GetContentTypes(w http.ResponseWriter, r *http.Request) {
	contentTypes, err := api.GetContentTypes()

	if err != nil {
		utils.PrintError(w, err, false)
		return
	}

	output := utils.ConstructOutput()
	output["data"] = contentTypes

	json, err := json.Marshal(output)

	if err != nil {
		utils.PrintError(w, err, true)
		return
	}

	fmt.Fprint(w, string(json))
}
