package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	api_structs "github.com/Dobefu/csb/cmd/api/structs"
	api_utils "github.com/Dobefu/csb/cmd/api/utils"
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
	"github.com/Dobefu/csb/cmd/server/utils"
	"github.com/stretchr/testify/assert"
)

type MockRows struct {
	nextCalls int
	scanError error
}

func (m *MockRows) Next() bool {
	m.nextCalls++
	return m.nextCalls == 1
}

func (m *MockRows) Scan(dest ...interface{}) error {
	if m.scanError != nil {
		return m.scanError
	}

	*dest[0].(*string) = "test-uid"
	*dest[1].(*string) = "test-locale"
	*dest[2].(*string) = "test-url"
	*dest[3].(*time.Time) = time.Unix(0, 0)

	return nil
}

var mockQueryRows func(table string, columns []string, where []db_structs.QueryWhere) (queryRows, error)

func setupGetSitemapData() (*httptest.ResponseRecorder, func()) {
	rr := httptest.NewRecorder()

	queryQueryRows = func(table string, columns []string, where []db_structs.QueryWhere) (queryRows, error) {
		return mockQueryRows(table, columns, where)
	}

	apiUtilsGetAltLocales = func(entry structs.Route, includeSitemapExcluded bool) ([]api_structs.AltLocale, error) {
		return []api_structs.AltLocale{}, nil
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
		apiUtilsGetAltLocales = api_utils.GetAltLocales
		utilsPrintError = utils.PrintError
		utilsConstructOutput = utils.ConstructOutput

		mockQueryRows = func(_ string, _ []string, _ []db_structs.QueryWhere) (queryRows, error) {
			return &MockRows{}, nil
		}
	}
}

func TestGetSitemapDataSuccess(t *testing.T) {
	rr, teardown := setupGetSitemapData()
	defer teardown()

	mockQueryRows = func(_ string, _ []string, _ []db_structs.QueryWhere) (queryRows, error) {
		return &MockRows{}, nil
	}

	req, _ := http.NewRequest("GET", "/sitemap-data", nil)

	GetSitemapData(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	data, ok := response["data"].(map[string]interface{})
	assert.True(t, ok)
	assert.Contains(t, data, "test-uid")

	entry := data["test-uid"].(map[string]interface{})
	assert.Equal(t, "test-uid", entry["uid"])
	assert.Equal(t, "test-locale", entry["locale"])
	assert.Equal(t, "test-url", entry["url"])
	assert.NotEmpty(t, entry["updated_at"])
	assert.Empty(t, entry["alt_locales"])
}

func TestGetSitemapDataErrJsonMarshal(t *testing.T) {
	rr, teardown := setupGetSitemapData()
	defer teardown()

	mockQueryRows = func(_ string, _ []string, _ []db_structs.QueryWhere) (queryRows, error) {
		return &MockRows{}, nil
	}

	utilsConstructOutput = func() map[string]map[string]interface{} {
		return map[string]map[string]interface{}{
			"data":    {},
			"invalid": {"key": make(chan int)},
			"error":   nil,
		}
	}

	req, _ := http.NewRequest("GET", "/sitemap-data", nil)

	GetSitemapData(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.JSONEq(t, `{"data":null,"error":"json: unsupported type: chan int"}`, rr.Body.String())
}

func TestGetSitemapDataErrGetEntriesErrQueryRows(t *testing.T) {
	rr, teardown := setupGetSitemapData()
	defer teardown()

	mockQueryRows = func(_ string, _ []string, _ []db_structs.QueryWhere) (queryRows, error) {
		return nil, errors.New("cannot query rows")
	}

	req, _ := http.NewRequest("GET", "/sitemap-data", nil)

	GetSitemapData(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.JSONEq(t, `{"data":null,"error":"cannot query rows"}`, rr.Body.String())
}

func TestGetSitemapDataErrGetEntriesErrScan(t *testing.T) {
	rr, teardown := setupGetSitemapData()
	defer teardown()

	mockQueryRows = func(_ string, _ []string, _ []db_structs.QueryWhere) (queryRows, error) {
		return &MockRows{scanError: errors.New("scan error")}, nil
	}

	req, _ := http.NewRequest("GET", "/sitemap-data", nil)

	GetSitemapData(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.JSONEq(t, `{"data":null,"error":"scan error"}`, rr.Body.String())
}

func TestGetSitemapDataErrGetEntriesErrAltLocales(t *testing.T) {
	rr, teardown := setupGetSitemapData()
	defer teardown()

	apiUtilsGetAltLocales = func(entry structs.Route, includeSitemapExcluded bool) ([]api_structs.AltLocale, error) {
		return nil, errors.New("cannot get alt locales")
	}

	req, _ := http.NewRequest("GET", "/sitemap-data", nil)

	GetSitemapData(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.JSONEq(t, `{"data":null,"error":"cannot get alt locales"}`, rr.Body.String())
}
