package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Dobefu/csb/cmd/server/utils"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	tests := []struct {
		name           string
		path           string
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:           "Home page",
			path:           "/",
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"data": map[string]interface{}{
					"api_endpoints": []interface{}{"/api"},
				},
				"error": nil,
			},
		},
		{
			name:           "Not found",
			path:           "/nonexistent",
			expectedStatus: http.StatusNotFound,
			expectedBody:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.path, nil)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				Index(w, r, "/api")
			})

			handler.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			if tt.expectedBody != nil {
				var got map[string]interface{}
				err = json.Unmarshal(rr.Body.Bytes(), &got)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, got)
			}
		})
	}
}

func TestIndexErrorHandling(t *testing.T) {
	utilsConstructOutput = func() map[string]map[string]interface{} {
		return map[string]map[string]interface{}{
			"data": {
				"invalid": make(chan int),
			},
			"error": nil,
		}
	}

	defer func() { utilsConstructOutput = utils.ConstructOutput }()

	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Index(w, r, "/api")
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
