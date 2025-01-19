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

func setupGetGlobalFields() (*httptest.ResponseRecorder, func()) {
	rr := httptest.NewRecorder()

	apiGetGlobalFields = func() (map[string]interface{}, error) {
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
		apiGetGlobalFields = api.GetGlobalFields
		utilsPrintError = utils.PrintError
	}
}

func TestGetGlobalFieldsSuccess(t *testing.T) {
	rr, teardown := setupGetGlobalFields()
	defer teardown()

	req, _ := http.NewRequest("GET", "/global-fields", nil)

	GetGlobalFields(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"data":{},"error":null}`, rr.Body.String())
}

func TestGetGlobalFieldsErrApiGet(t *testing.T) {
	rr, teardown := setupGetGlobalFields()
	defer teardown()

	apiGetGlobalFields = func() (map[string]interface{}, error) {
		return nil, errors.New("cannot get global fields")
	}

	req, _ := http.NewRequest("GET", "/global-fields", nil)

	GetGlobalFields(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.JSONEq(t, `{"data":null,"error":"cannot get global fields"}`, rr.Body.String())
}

func TestGetGlobalFieldsErrJsonMarshal(t *testing.T) {
	rr, teardown := setupGetGlobalFields()
	defer teardown()

	apiGetGlobalFields = func() (map[string]interface{}, error) {
		return map[string]interface{}{"invalid": make(chan int)}, nil
	}

	req, _ := http.NewRequest("GET", "/global-fields", nil)

	GetGlobalFields(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.JSONEq(t, `{"data":null,"error":"json: unsupported type: chan int"}`, rr.Body.String())
}
