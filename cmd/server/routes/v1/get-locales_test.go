package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLocales(t *testing.T) {
	originalGetLocales := apiGetLocales
	originalConstructOutput := utilsConstructOutput
	originalPrintError := utilsPrintError

	defer func() {
		apiGetLocales = originalGetLocales
		utilsConstructOutput = originalConstructOutput
		utilsPrintError = originalPrintError
	}()

	tests := []struct {
		name           string
		mockGetLocales func() (map[string]interface{}, error)
		mockConstruct  func() map[string]map[string]interface{}
		mockPrintError func(w http.ResponseWriter, err error, internal bool)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Success",
			mockGetLocales: func() (map[string]interface{}, error) {
				return map[string]interface{}{
					"locales": []interface{}{
						map[string]interface{}{"code": "en-US", "name": "English - United States"},
						map[string]interface{}{"code": "es-ES", "name": "Dutch - Netherlands"},
					},
				}, nil
			},
			mockConstruct: func() map[string]map[string]interface{} {
				return map[string]map[string]interface{}{"data": nil, "error": nil}
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"data":{"locales":[{"code":"en-US","name":"English - United States"},{"code":"es-ES","name":"Dutch - Netherlands"}]},"error":null}`,
		},
		{
			name: "API Error",
			mockGetLocales: func() (map[string]interface{}, error) {
				return nil, errors.New("API error")
			},
			mockPrintError: func(w http.ResponseWriter, err error, internal bool) {
				assert.False(t, internal)
				w.WriteHeader(http.StatusInternalServerError)
				_ = json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"API error"}`,
		},
		{
			name: "Marshal Error",
			mockGetLocales: func() (map[string]interface{}, error) {
				return map[string]interface{}{"locales": make(chan int)}, nil
			},
			mockConstruct: func() map[string]map[string]interface{} {
				return map[string]map[string]interface{}{"data": nil, "error": nil}
			},
			mockPrintError: func(w http.ResponseWriter, err error, internal bool) {
				assert.True(t, internal)
				w.WriteHeader(http.StatusInternalServerError)
				_ = json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"json: unsupported type: chan int"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiGetLocales = tt.mockGetLocales

			if tt.mockConstruct != nil {
				utilsConstructOutput = tt.mockConstruct
			}

			if tt.mockPrintError != nil {
				utilsPrintError = tt.mockPrintError
			}

			req, err := http.NewRequest("GET", "/locales", nil)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()

			GetLocales(rr, req)
			assert.Equal(t, tt.expectedStatus, rr.Code)
			assert.JSONEq(t, tt.expectedBody, rr.Body.String())
		})
	}
}
