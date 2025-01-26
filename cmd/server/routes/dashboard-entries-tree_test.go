package routes

import (
	"bytes"
	"crypto/rsa"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/fstest"

	"github.com/Dobefu/csb/cmd/cs_sdk"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

var mockClient = &MockClient{
	DoFunc: func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(
				`{"signing-key":"test-key"}`,
			)),
		}, nil
	},
}

func setupDashboardEntriesTreeTest() (*httptest.ResponseRecorder, func()) {
	rr := httptest.NewRecorder()

	originalClient := httpClient
	httpClient = mockClient

	csSdkGetUrl = func(useManagementToken bool) string {
		return ""
	}

	jwtParse = func(tokenString string, keyFunc jwt.Keyfunc, options ...jwt.ParserOption) (*jwt.Token, error) {
		return &jwt.Token{Valid: true}, nil
	}

	jwtParseRSAPublicKeyFromPEM = func(key []byte) (*rsa.PublicKey, error) {
		return &rsa.PublicKey{}, nil
	}

	return rr, func() {
		getFs = func() FS { return content }
		httpClient = originalClient
		csSdkGetUrl = cs_sdk.GetUrl
		jwtParse = jwt.Parse
		jwtParseRSAPublicKeyFromPEM = jwt.ParseRSAPublicKeyFromPEM
	}
}

func TestDashboardEntriesTreeSuccess(t *testing.T) {
	rr, cleanup := setupDashboardEntriesTreeTest()
	defer cleanup()

	req, _ := http.NewRequest("GET", "/?app-token=test-token", nil)
	DashboardEntriesTree(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.NotEmpty(t, rr.Body.String())
}

func TestDashboardEntriesTreeErrNoTemplate(t *testing.T) {
	rr, cleanup := setupDashboardEntriesTreeTest()
	defer cleanup()

	getFs = func() FS { return fstest.MapFS{} }

	req, _ := http.NewRequest("GET", "/?app-token=test-token", nil)
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

	req, _ := http.NewRequest("GET", "/?app-token=test-token", nil)
	DashboardEntriesTree(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "", rr.Body.String())
}
