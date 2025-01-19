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

func setupGetContentTypesTest() (*httptest.ResponseRecorder, func()) {
	utilsPrintError = func(w http.ResponseWriter, err error, logError bool) {
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"data": nil, "error": err.Error(),
			},
		)
	}

	rr := httptest.NewRecorder()

	cleanup := func() {
		apiGetContentTypes = api.GetContentTypes
		utilsPrintError = utils.PrintError
	}

	return rr, cleanup
}

func TestGetContentTypesSuccess(t *testing.T) {
	apiGetContentTypes = func() (map[string]interface{}, error) {
		return map[string]interface{}{"name": "blog", "fields": []string{"title", "content"}}, nil
	}

	rr, cleanup := setupGetContentTypesTest()
	defer cleanup()

	req, _ := http.NewRequest("GET", "/content-types", nil)
	GetContentTypes(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"data":{"name":"blog","fields":["title","content"]},"error":null}`, rr.Body.String())
}

func TestGetContentTypesErrApi(t *testing.T) {
	apiGetContentTypes = func() (map[string]interface{}, error) {
		return nil, errors.New("failed to get content types")
	}

	rr, cleanup := setupGetContentTypesTest()
	defer cleanup()

	req, _ := http.NewRequest("GET", "/content-type", nil)
	GetContentTypes(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.JSONEq(t, `{"data":null,"error":"failed to get content types"}`, rr.Body.String())
}

func TestGetContentTypesErrJsonMarshal(t *testing.T) {
	apiGetContentTypes = func() (map[string]interface{}, error) {
		return map[string]interface{}{"data": make(chan int), "error": nil}, nil
	}

	rr, cleanup := setupGetContentTypesTest()
	defer cleanup()

	req, _ := http.NewRequest("GET", "/content-type", nil)
	GetContentTypes(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.JSONEq(t, `{"data":null,"error":"json: unsupported type: chan int"}`, rr.Body.String())
}
