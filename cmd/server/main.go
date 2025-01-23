package server

import (
	"fmt"
	"net/http"

	"github.com/Dobefu/csb/cmd/logger"
)

func Start(port uint) error {
	url := fmt.Sprintf(":%d", port)
	mux := http.NewServeMux()
	apiPath := "/api/v1"

	HandleRoutes(mux, apiPath)

	logger.Info("Starting server on %s", url)

	err := listenAndServe(url, mux)

	if err != nil {
		return err
	}

	return nil
}
