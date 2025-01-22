package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Dobefu/csb/cmd/api"
	"github.com/Dobefu/csb/cmd/server/utils"
	"github.com/Dobefu/csb/cmd/server/validation"
)

var validationCheckRequiredQueryParams = validation.CheckRequiredQueryParams
var apiGetContentType = api.GetContentType

func GetContentType(w http.ResponseWriter, r *http.Request) {
	params, err := validationCheckRequiredQueryParams(
		r,
		"content_type",
	)

	if err != nil {
		utilsPrintError(w, err, false)
		return
	}

	contentTypeName, hasContentTypeName := params["content_type"]

	if !hasContentTypeName {
		utilsPrintError(w, errors.New("no content type name found"), false)
		return
	}

	contentType, err := apiGetContentType(contentTypeName.(string))

	if err != nil {
		utilsPrintError(w, err, false)
		return
	}

	output := utils.ConstructOutput()
	output["data"] = contentType

	json, err := json.Marshal(output)

	if err != nil {
		utilsPrintError(w, err, true)
		return
	}

	fmt.Fprint(w, string(json))
}
