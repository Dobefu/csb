package middleware

import (
	"errors"
	"net/http"
	"os"

	"github.com/Dobefu/csb/cmd/logger"
	"github.com/Dobefu/csb/cmd/server/utils"
)

func init() {
	authDebug := os.Getenv("DEBUG_AUTH_BYPASS")

	if authDebug != "" {
		logger.Warning("DEBUG_AUTH_BYPASS is set. Running without any auth token checks")
	}
}

func RequireDeliveryToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authDebug := os.Getenv("DEBUG_AUTH_BYPASS")

		// If auth debugging is enabled, skip this middleware.
		if authDebug != "" {
			next.ServeHTTP(w, r)
			return
		}

		token := os.Getenv("CS_DELIVERY_TOKEN")

		if token == "" {
			logger.Error("CS_DELIVERY_TOKEN is not set")
		}

		authToken := r.Header.Get("Authorization")

		if authToken != token {
			http.Error(w, "", http.StatusForbidden)
			utils.PrintError(w, errors.New("Invalid authorization token"), false)
			return
		}

		next.ServeHTTP(w, r)
	})
}
