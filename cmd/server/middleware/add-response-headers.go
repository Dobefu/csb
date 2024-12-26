package middleware

import (
	"net/http"
	"os"

	"github.com/Dobefu/csb/cmd/logger"
)

func init() {
	authDebug := os.Getenv("DEBUG_AUTH_BYPASS")

	if authDebug != "" {
		logger.Warning("DEBUG_AUTH_BYPASS is set. Running without any auth token checks")
	}
}

func AddResponseHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
