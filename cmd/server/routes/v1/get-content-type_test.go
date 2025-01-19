package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Dobefu/csb/cmd/api"
	"github.com/Dobefu/csb/cmd/server/utils"
	"github.com/Dobefu/csb/cmd/server/validation"
	"github.com/stretchr/testify/assert"
)

func setupGetContentTypeTest() (*httptest.ResponseRecorder, func()) {
	validationCheckRequiredQueryParams = func(r *http.Request, params ...string) (map[string]interface{}, error) {
		return map[string]interface{}{"content_type": "blog"}, nil
	}

	apiGetContentType = func(contentTypeName string) (map[string]interface{}, error) {
		return map[string]interface{}{"name": "blog", "fields": []string{"title", "content"}}, nil
	}

	utilsPrintError = func(w http.ResponseWriter, err error, internal bool) {
		if internal {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"data": nil, "error": err.Error()})
	}

	rr := httptest.NewRecorder()

	cleanup := func() {
		validationCheckRequiredQueryParams = validation.CheckRequiredQueryParams
		apiGetContentType = api.GetContentType
		utilsPrintError = utils.PrintError
	}

	return rr, cleanup
}

func TestGetContentTypeSuccess(t *testing.T) {
	rr, cleanup := setupGetContentTypeTest()
	defer cleanup()

	req, _ := http.NewRequest("GET", "/content-type?content_type=blog", nil)
	GetContentType(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"data":{"name":"blog","fields":["title","content"]},"error":null}`, rr.Body.String())
}

func TestGetContentTypeMissingQueryParameter(t *testing.T) {
	rr, cleanup := setupGetContentTypeTest()
	defer cleanup()

	validationCheckRequiredQueryParams = func(r *http.Request, params ...string) (map[string]interface{}, error) {
		return nil, errors.New("missing required parameter: content_type")
	}

	req, _ := http.NewRequest("GET", "/content-type", nil)
	GetContentType(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.JSONEq(t, `{"data":null,"error":"missing required parameter: content_type"}`, rr.Body.String())
}

func TestGetContentTypeNoContentTypeName(t *testing.T) {
	rr, cleanup := setupGetContentTypeTest()
	defer cleanup()

	validationCheckRequiredQueryParams = func(r *http.Request, params ...string) (map[string]interface{}, error) {
		return map[string]interface{}{}, nil
	}

	req, _ := http.NewRequest("GET", "/content-type?content_type=", nil)
	GetContentType(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.JSONEq(t, `{"data":null,"error":"no content type name found"}`, rr.Body.String())
}

func TestGetContentTypeAPIError(t *testing.T) {
	rr, cleanup := setupGetContentTypeTest()
	defer cleanup()

	apiGetContentType = func(contentTypeName string) (map[string]interface{}, error) {
		return nil, errors.New("API error")
	}

	req, _ := http.NewRequest("GET", "/content-type?content_type=blog", nil)
	GetContentType(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.JSONEq(t, `{"data":null,"error":"API error"}`, rr.Body.String())
}

func TestGetContentTypeJSONMarshalError(t *testing.T) {
	rr, cleanup := setupGetContentTypeTest()
	defer cleanup()

	apiGetContentType = func(contentTypeName string) (map[string]interface{}, error) {
		return map[string]interface{}{"unmarshalable": make(chan int)}, nil
	}

	req, _ := http.NewRequest("GET", "/content-type?content_type=blog", nil)
	GetContentType(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Contains(t, rr.Body.String(), "json: unsupported type: chan int")
}
