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
	apiRoute(mux, apiPath, "/content-type", "GET", v1.GetContentType)
	apiRoute(mux, apiPath, "/global-fields", "GET", v1.GetGlobalFields)
	apiRoute(mux, apiPath, "/locales", "GET", v1.GetLocales)
	apiRoute(mux, apiPath, "/translations", "GET", v1.GetTranslations)
	apiRoute(mux, apiPath, "/sitemap-data", "GET", v1.GetSitemapData)
	apiRoute(mux, apiPath, "/sync", "POST", v1.Sync)
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
		middleware.RequireDeliveryToken(
			middleware.AddResponseHeaders(
				http.HandlerFunc(handler),
			),
		),
	)
}
