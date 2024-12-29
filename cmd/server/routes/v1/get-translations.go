package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dobefu/csb/cmd/api"
	"github.com/Dobefu/csb/cmd/server/utils"
	"github.com/Dobefu/csb/cmd/server/validation"
)

func GetTranslations(w http.ResponseWriter, r *http.Request) {
	params, err := validation.CheckRequiredQueryParams(
		r,
		"locale",
	)

	if err != nil {
		utils.PrintError(w, err, false)
		return
	}

	locale := params["locale"].(string)

	translations, err := api.GetTranslations(locale)

	if err != nil {
		utils.PrintError(w, err, false)
		return
	}

	output := utils.ConstructOutput()
	output["data"] = translations

	json, err := json.Marshal(output)

	if err != nil {
		utils.PrintError(w, err, true)
		return
	}

	fmt.Fprint(w, string(json))
}
