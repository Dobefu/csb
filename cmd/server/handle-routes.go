package server

import (
	"fmt"
	"net/http"

	"github.com/Dobefu/csb/cmd/server/routes"
	v1 "github.com/Dobefu/csb/cmd/server/routes/v1"
)

func HandleRoutes(mux *http.ServeMux) {
	apiPath := "/api/v1"

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		routes.Index(w, r, apiPath)
	})

	mux.HandleFunc(fmt.Sprintf("GET %s/get-entry-by-url", apiPath), v1.GetEntryByUrl)
	mux.HandleFunc(fmt.Sprintf("GET %s/get-entry-by-uid", apiPath), v1.GetEntryByUid)
}
