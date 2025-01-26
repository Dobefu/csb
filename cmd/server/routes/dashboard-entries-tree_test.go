package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/fstest"

	"github.com/Dobefu/csb/cmd/cs_sdk"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func setupDashboardEntriesTreeTest() (*httptest.ResponseRecorder, func()) {
	rr := httptest.NewRecorder()

	csSdkGetUrl = func(useManagementToken bool) string {
		return ""
	}

	return rr, func() {
		getFs = func() FS { return content }
		jwtParseRSAPublicKeyFromPEM = jwt.ParseRSAPublicKeyFromPEM
		csSdkGetUrl = cs_sdk.GetUrl
	}
}

func TestDashboardEntriesTreeSuccess(t *testing.T) {
	rr, cleanup := setupDashboardEntriesTreeTest()
	defer cleanup()

	req, _ := http.NewRequest("GET", "/", nil)
	DashboardEntriesTree(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.NotEmpty(t, rr.Body.String())
}

func TestDashboardEntriesTreeErrNoTemplate(t *testing.T) {
	rr, cleanup := setupDashboardEntriesTreeTest()
	defer cleanup()

	getFs = func() FS { return fstest.MapFS{} }

	req, _ := http.NewRequest("GET", "/", nil)
	DashboardEntriesTree(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "", rr.Body.String())
}

func TestDashboardEntriesTreeErrInvalidTemplate(t *testing.T) {
	rr, cleanup := setupDashboardEntriesTreeTest()
	defer cleanup()

	getFs = func() FS {
		return fstest.MapFS{
			"templates/dashboard-entries-tree.html.tmpl": {
				Data: []byte("{{ .Bogus }}"),
			},
		}
	}

	req, _ := http.NewRequest("GET", "/", nil)
	DashboardEntriesTree(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "", rr.Body.String())
}
