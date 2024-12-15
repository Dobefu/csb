package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dobefu/csb/cmd/api"
	"github.com/Dobefu/csb/cmd/server/utils"
)

func GetContentTypes(w http.ResponseWriter, r *http.Request) {
	output, err := api.GetContentTypes()

	if err != nil {
		utils.PrintError(w, err, false)
		return
	}

	json, err := json.Marshal(output["content_types"])

	if err != nil {
		utils.PrintError(w, err, true)
		return
	}

	fmt.Fprint(w, string(json))
}
