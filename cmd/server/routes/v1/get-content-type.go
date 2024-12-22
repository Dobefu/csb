package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dobefu/csb/cmd/api"
	"github.com/Dobefu/csb/cmd/server/utils"
	"github.com/Dobefu/csb/cmd/server/validation"
)

func GetContentType(w http.ResponseWriter, r *http.Request) {
	params, err := validation.CheckRequiredQueryParams(
		r,
		"content_type",
	)

	if err != nil {
		utils.PrintError(w, err, false)
		return
	}

	contentTypeName := params["content_type"].(string)

	contentType, err := api.GetContentType(contentTypeName)

	if err != nil {
		utils.PrintError(w, err, false)
		return
	}

	output := utils.ConstructOutput()
	output["data"] = contentType

	json, err := json.Marshal(output)

	if err != nil {
		utils.PrintError(w, err, true)
		return
	}

	fmt.Fprint(w, string(json))
}
