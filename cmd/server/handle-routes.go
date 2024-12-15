package server

import (
	"fmt"
	"net/http"

	"github.com/Dobefu/csb/cmd/server/middleware"
	"github.com/Dobefu/csb/cmd/server/routes"
	v1 "github.com/Dobefu/csb/cmd/server/routes/v1"
)

func HandleRoutes(mux *http.ServeMux, apiPath string) {
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		routes.Index(w, r, apiPath)
	})

	apiRoute(mux, apiPath, "/get-entry-by-url", "GET", v1.GetEntryByUrl)
	apiRoute(mux, apiPath, "/get-entry-by-uid", "GET", v1.GetEntryByUid)
	apiRoute(mux, apiPath, "/content-types", "GET", v1.GetContentTypes)
}

func apiRoute(
	mux *http.ServeMux,
	apiPath string,
	path string,
	method string,
	handler func(w http.ResponseWriter, r *http.Request),
) {
	fullPath := fmt.Sprintf("%s %s%s", method, apiPath, path)
	mux.Handle(
		fullPath,
		middleware.RequireDeliveryToken(http.HandlerFunc(handler)),
	)
}
