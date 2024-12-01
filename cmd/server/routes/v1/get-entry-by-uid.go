package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dobefu/csb/cmd/api"
	"github.com/Dobefu/csb/cmd/logger"
	"github.com/Dobefu/csb/cmd/server/validation"
)

func GetEntryByUid(w http.ResponseWriter, r *http.Request) {
	params, err := validation.CheckRequiredQueryParams(
		r,
		"uid",
		"locale",
	)

	if err != nil {
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
		return
	}

	uid := params["uid"].(string)
	locale := params["locale"].(string)

	entry, err := api.GetEntryByUid(uid, locale, false)

	if err != nil {
		fmt.Fprintf(w, `{"data": null, "error": "%s"}`, err.Error())
		return
	}

	output := map[string]interface{}{
		"data": map[string]interface{}{
			"entry": entry,
		},
		"error": nil,
	}

	json, err := json.Marshal(output)

	if err != nil {
		logger.Error(err.Error())
		fmt.Fprintf(w, `{"data": null, "error": "%s"}`, err.Error())
		return
	}

	fmt.Fprint(w, string(json))
}
