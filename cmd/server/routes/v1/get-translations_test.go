package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Dobefu/csb/cmd/api"
	"github.com/Dobefu/csb/cmd/server/utils"
	"github.com/stretchr/testify/assert"
)

func setupGetTranslations() (*httptest.ResponseRecorder, func()) {
	rr := httptest.NewRecorder()

	apiGetTranslations = func(locale string) (map[string]interface{}, error) {
		return map[string]interface{}{}, nil
	}

	utilsPrintError = func(w http.ResponseWriter, err error, internal bool) {
		if internal {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}

		_ = json.NewEncoder(w).Encode(map[string]interface{}{"data": nil, "error": err.Error()})
	}

	return rr, func() {
		apiGetTranslations = api.GetTranslations
		utilsPrintError = utils.PrintError
	}
}

func TestGetTranslationsSuccess(t *testing.T) {
	rr, teardown := setupGetTranslations()
	defer teardown()

	req, _ := http.NewRequest("GET", "/translations?locale=en", nil)

	GetTranslations(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"data":{},"error":null}`, rr.Body.String())
}

func TestGetTranslationsErrMissingParams(t *testing.T) {
	rr, teardown := setupGetTranslations()
	defer teardown()

	req, _ := http.NewRequest("GET", "/translations", nil)

	GetTranslations(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.JSONEq(t, `{"data":null,"error":"missing required query params: (locale)"}`, rr.Body.String())
}

func TestGetTranslationsErrApiGet(t *testing.T) {
	rr, teardown := setupGetTranslations()
	defer teardown()

	apiGetTranslations = func(locale string) (map[string]interface{}, error) {
		return nil, errors.New("cannot get translations")
	}

	req, _ := http.NewRequest("GET", "/translations?locale=en", nil)

	GetTranslations(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.JSONEq(t, `{"data":null,"error":"cannot get translations"}`, rr.Body.String())
}

func TestGetTranslationsErrJsonMarshal(t *testing.T) {
	rr, teardown := setupGetTranslations()
	defer teardown()

	apiGetTranslations = func(locale string) (map[string]interface{}, error) {
		return map[string]interface{}{"invalid": make(chan int)}, nil
	}

	req, _ := http.NewRequest("GET", "/translations?locale=en", nil)

	GetTranslations(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.JSONEq(t, `{"data":null,"error":"json: unsupported type: chan int"}`, rr.Body.String())
}
