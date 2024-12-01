package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dobefu/csb/cmd/logger"
)

func Index(w http.ResponseWriter, r *http.Request, apiPath string) {
	// If not on the homepage, return a 404.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	output := map[string]interface{}{
		"data": map[string]interface{}{
			"api_endpoints": []string{apiPath},
		},
		"error": nil,
	}

	json, err := json.Marshal(output)

	if err != nil {
		logger.Error(err.Error())
		fmt.Fprintf(w, `{"error": %s}`, err.Error())
		return
	}

	fmt.Fprint(w, string(json))
}
