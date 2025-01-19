package v1

import (
	"fmt"
	"net/http"

	"github.com/Dobefu/csb/cmd/cs_sdk/functions"
)

var functionsSync = functions.Sync

func Sync(w http.ResponseWriter, r *http.Request) {
	err := functionsSync(false)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utilsPrintError(w, err, true)
		return
	}

	fmt.Fprint(w, string(`{"error": null}`))
}
