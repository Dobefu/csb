package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Dobefu/csb/cmd/cs_sdk/functions"
	"github.com/Dobefu/csb/cmd/server/utils"
	"github.com/stretchr/testify/assert"
)

func setupSync() (*httptest.ResponseRecorder, func()) {
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
		functionsSync = functions.Sync
		utilsPrintError = utils.PrintError
	}
}

func TestSyncSuccess(t *testing.T) {
	rr, teardown := setupSync()
	defer teardown()

	functionsSync = func(reset bool) error {
		return nil
	}

	req, _ := http.NewRequest("GET", "/sync", nil)

	Sync(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"error":null}`, rr.Body.String())
}

func TestSyncErrApiGet(t *testing.T) {
	rr, teardown := setupSync()
	defer teardown()

	functionsSync = func(reset bool) error {
		return errors.New("cannot sync")
	}

	req, _ := http.NewRequest("GET", "/sync", nil)

	Sync(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.JSONEq(t, `{"data":null,"error":"cannot sync"}`, rr.Body.String())
}
