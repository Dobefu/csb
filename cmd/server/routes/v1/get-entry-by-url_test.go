package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Dobefu/csb/cmd/api"
	api_structs "github.com/Dobefu/csb/cmd/api/structs"
	cs_api "github.com/Dobefu/csb/cmd/cs_sdk/api"
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/Dobefu/csb/cmd/server/utils"
	"github.com/stretchr/testify/assert"
)

func setupGetEntryByUrl() (*httptest.ResponseRecorder, func()) {
	rr := httptest.NewRecorder()

	apiGetEntryByUrl = func(url string, locale string, includeUnpublished bool) (structs.Route, error) {
		return structs.Route{}, nil
	}

	csApiGetEntryWithMetadata = func(route structs.Route) (interface{}, []api_structs.AltLocale, []structs.Route, error) {
		return map[string]interface{}{}, []api_structs.AltLocale{}, []structs.Route{}, nil
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
		apiGetEntryByUrl = api.GetEntryByUrl
		csApiGetEntryWithMetadata = cs_api.GetEntryWithMetadata
		utilsPrintError = utils.PrintError
		utilsConstructEntryOutput = utils.ConstructEntryOutput
	}
}

func TestGetEntryByUrlSuccess(t *testing.T) {
	rr, teardown := setupGetEntryByUrl()
	defer teardown()

	req, _ := http.NewRequest("GET", "/entry-by-url?url=test&locale=en", nil)

	GetEntryByUrl(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"data":{"alt_locales":[],"breadcrumbs":[],"entry":{}},"error":null}`, rr.Body.String())
}

func TestGetEntryByUrlErrMissingParams(t *testing.T) {
	rr, teardown := setupGetEntryByUrl()
	defer teardown()

	req, _ := http.NewRequest("GET", "/entry-by-url", nil)

	GetEntryByUrl(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.JSONEq(t, `{"data":null,"error":"missing required query params: (url, locale)"}`, rr.Body.String())
}

func TestGetEntryByUrlErrApiGetEntry(t *testing.T) {
	rr, teardown := setupGetEntryByUrl()
	defer teardown()

	apiGetEntryByUrl = func(url string, locale string, includeUnpublished bool) (structs.Route, error) {
		return structs.Route{}, errors.New("cannot get the entry")
	}

	req, _ := http.NewRequest("GET", "/entry-by-url?url=test&locale=en", nil)

	GetEntryByUrl(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.JSONEq(t, `{"data":null,"error":"cannot get the entry"}`, rr.Body.String())
}

func TestGetEntryByUrlErrApiGetMetadata(t *testing.T) {
	rr, teardown := setupGetEntryByUrl()
	defer teardown()

	csApiGetEntryWithMetadata = func(route structs.Route) (interface{}, []api_structs.AltLocale, []structs.Route, error) {
		return nil, []api_structs.AltLocale{}, []structs.Route{}, errors.New("cannot get the entry metadata")
	}

	req, _ := http.NewRequest("GET", "/entry-by-url?url=test&locale=en", nil)

	GetEntryByUrl(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.JSONEq(t, `{"data":null,"error":"cannot get the entry metadata"}`, rr.Body.String())
}

func TestGetEntryByUrlErrJson(t *testing.T) {
	rr, teardown := setupGetEntryByUrl()
	defer teardown()

	utilsConstructEntryOutput = func(
		entry interface{},
		altLocales []api_structs.AltLocale,
		breadcrumbs []structs.Route,
	) map[string]map[string]interface{} {
		return map[string]map[string]interface{}{
			"entry": {"invalid": make(chan int)},
		}
	}

	req, _ := http.NewRequest("GET", "/entry-by-url?url=test&locale=en", nil)

	GetEntryByUrl(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.JSONEq(t, `{"data":null,"error":"json: unsupported type: chan int"}`, rr.Body.String())
}
