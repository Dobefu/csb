package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dobefu/csb/cmd/api"
	"github.com/Dobefu/csb/cmd/logger"
	"github.com/Dobefu/csb/cmd/server/utils"
	"github.com/Dobefu/csb/cmd/server/validation"
)

func GetEntryByUrl(w http.ResponseWriter, r *http.Request) {
	params, err := validation.CheckRequiredQueryParams(
		r,
		"url",
		"locale",
	)

	if err != nil {
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
		return
	}

	url := params["url"].(string)
	locale := params["locale"].(string)

	entry, err := api.GetEntryByUrl(url, locale, false)

	if err != nil {
		utils.PrintError(w, err)
		return
	}

	output := utils.ConstructOutput()
	output["data"]["entry"] = entry

	json, err := json.Marshal(output)

	if err != nil {
		logger.Error(err.Error())
		utils.PrintError(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}
