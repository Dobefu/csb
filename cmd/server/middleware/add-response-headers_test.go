package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupAddResponseHeaders() *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()

	utilsPrintError = func(w http.ResponseWriter, err error, internal bool) {
		if internal {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}

		_ = json.NewEncoder(w).Encode(map[string]interface{}{"data": nil, "error": err.Error()})
	}

	return rr
}

func TestAddResponseHeadersSuccess(t *testing.T) {
	rr := setupAddResponseHeaders()
	req, _ := http.NewRequest("GET", "/", nil)

	endpoint := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"error":null}`)
	})

	AddResponseHeaders(endpoint).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"error":null}`, rr.Body.String())
}
