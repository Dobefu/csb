package server

import (
	"net/http"

	"github.com/Dobefu/csb/cmd/server/routes"
)

func HandleRoutes(mux *http.ServeMux) {
	apiPath := "/api/v1"

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		routes.Index(w, r, apiPath)
	})
}
