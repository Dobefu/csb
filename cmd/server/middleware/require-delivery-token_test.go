package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Dobefu/csb/cmd/server/utils"
	"github.com/stretchr/testify/assert"
)

func setupRequireDeliveryToken() (*httptest.ResponseRecorder, func()) {
	rr := httptest.NewRecorder()

	utilsPrintError = func(w http.ResponseWriter, err error, internal bool) {
		if internal {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}

		_ = json.NewEncoder(w).Encode(map[string]interface{}{"data": nil, "error": err.Error()})
	}

	return rr, func() {
		utilsPrintError = utils.PrintError

		os.Unsetenv("CS_DELIVERY_TOKEN")
		os.Unsetenv("DEBUG_AUTH_BYPASS")
	}
}

func TestRequireDeliveryTokenSuccess(t *testing.T) {
	rr, cleanup := setupRequireDeliveryToken()
	defer cleanup()

	os.Setenv("CS_DELIVERY_TOKEN", "testing")

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "testing")

	endpoint := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"error":null}`)
	})

	RequireDeliveryToken(endpoint).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"error":null}`, rr.Body.String())
}

func TestRequireDeliveryTokenBypass(t *testing.T) {
	rr, cleanup := setupRequireDeliveryToken()
	defer cleanup()

	os.Setenv("DEBUG_AUTH_BYPASS", "1")
	os.Setenv("CS_DELIVERY_TOKEN", "testing")

	req, _ := http.NewRequest("GET", "/", nil)

	endpoint := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"error":null}`)
	})

	RequireDeliveryToken(endpoint).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"error":null}`, rr.Body.String())
}

func TestRequireDeliveryTokenErrAuthToken(t *testing.T) {
	rr, cleanup := setupRequireDeliveryToken()
	defer cleanup()

	os.Setenv("CS_DELIVERY_TOKEN", "testing")

	req, _ := http.NewRequest("GET", "/", nil)

	endpoint := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"error":null}`)
	})

	RequireDeliveryToken(endpoint).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusForbidden, rr.Code)
	assert.JSONEq(t, `{"data":null,"error":"Invalid authorization token"}`, rr.Body.String())
}

func TestRequireDeliveryTokenErrNoDeliveryToken(t *testing.T) {
	rr, cleanup := setupRequireDeliveryToken()
	defer cleanup()

	req, _ := http.NewRequest("GET", "/", nil)

	endpoint := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"error":null}`)
	})

	RequireDeliveryToken(endpoint).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "\n", rr.Body.String())
}
