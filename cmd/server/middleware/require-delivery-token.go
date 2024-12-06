package middleware

import (
	"errors"
	"net/http"
	"os"

	"github.com/Dobefu/csb/cmd/logger"
	"github.com/Dobefu/csb/cmd/server/utils"
)

func RequireDeliveryToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
