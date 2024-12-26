package v1

import (
	"fmt"
	"net/http"

	"github.com/Dobefu/csb/cmd/cs_sdk/functions"
	"github.com/Dobefu/csb/cmd/server/utils"
)

func Sync(w http.ResponseWriter, r *http.Request) {
	err := functions.Sync(false)

	if err != nil {
		utils.PrintError(w, err, false)
		return
	}

	fmt.Fprint(w, string(`{"error": null}`))
}
