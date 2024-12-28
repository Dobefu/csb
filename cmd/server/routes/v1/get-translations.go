package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dobefu/csb/cmd/api"
	"github.com/Dobefu/csb/cmd/server/utils"
)

func GetTranslations(w http.ResponseWriter, r *http.Request) {
	locales, err := api.GetTranslations()

	if err != nil {
		utils.PrintError(w, err, false)
		return
	}

	output := utils.ConstructOutput()
	output["data"] = locales

	json, err := json.Marshal(output)

	if err != nil {
		utils.PrintError(w, err, true)
		return
	}

	fmt.Fprint(w, string(json))
}
