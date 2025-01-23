package v1

import (
	"fmt"
	"net/http"
)

func Sync(w http.ResponseWriter, r *http.Request) {
	err := functionsSync(false)

	if err != nil {
		utilsPrintError(w, err, true)
		return
	}

	fmt.Fprint(w, string(`{"error": null}`))
}
