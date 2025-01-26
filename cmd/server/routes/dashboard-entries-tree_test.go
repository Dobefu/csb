package routes

import (
	"bytes"
	"crypto/rsa"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

type BrokenReader struct{}

func (br *BrokenReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("failed reading")
}

func (br *BrokenReader) Close() error {
	return fmt.Errorf("failed closing")
}

var mockClient = &MockClient{
	DoFunc: func(req *http.Request) (*http.Response, error) {
		if strings.Contains(req.URL.String(), "https://") {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body: io.NopCloser(bytes.NewBufferString(
					`{"signing-key":"test-key"}`,
				)),
			}, nil
		}

		if strings.Contains(req.URL.String(), "nil-body") {
			reader := BrokenReader{}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       &reader,
			}, nil
		}

		if strings.Contains(req.URL.String(), "json-err") {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body: io.NopCloser(bytes.NewBufferString(
					`{`,
				)),
			}, nil
		}

		if strings.Contains(req.URL.String(), "json-empty") {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body: io.NopCloser(bytes.NewBufferString(
					`{}`,
				)),
			}, nil
		}

		return nil, errors.New("invalid URL")
	},
}

func setupDashboardEntriesTreeTest() (*httptest.ResponseRecorder, func()) {
	rr := httptest.NewRecorder()

	originalClient := httpClient
	httpClient = mockClient

	csSdkGetUrl = func(useManagementToken bool) string {
		return "https://test.url"
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

	jwtParse = func(tokenString string, keyFunc jwt.Keyfunc, options ...jwt.ParserOption) (*jwt.Token, error) {
		key := &jwt.Token{Valid: true}
		_, err := keyFunc(key)
		assert.NoError(t, err)

		return key, nil
	}

	req, err := http.NewRequest("GET", "/?app-token=some.test.token", nil)
	assert.NoError(t, err)

	DashboardEntriesTree(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.NotEmpty(t, rr.Body.String())
}

func TestDashboardEntriesTreeErrParsePublicKeyFromPem(t *testing.T) {
	rr, cleanup := setupDashboardEntriesTreeTest()
	defer cleanup()

	jwtParse = func(tokenString string, keyFunc jwt.Keyfunc, options ...jwt.ParserOption) (*jwt.Token, error) {
		_, err := keyFunc(nil)
		assert.EqualError(t, err, "cannot parse token")

		return nil, errors.New("")
	}

	jwtParseRSAPublicKeyFromPEM = func(key []byte) (*rsa.PublicKey, error) {
		return nil, errors.New("cannot parse token")
	}

	req, err := http.NewRequest("GET", "/?app-token=test-token", nil)
	assert.NoError(t, err)

	DashboardEntriesTree(rr, req)
	assert.Equal(t, http.StatusForbidden, rr.Code)
}

func TestDashboardEntriesTreeErrInvalidPayload(t *testing.T) {
	rr, cleanup := setupDashboardEntriesTreeTest()
	defer cleanup()

	jwtParse = func(tokenString string, keyFunc jwt.Keyfunc, options ...jwt.ParserOption) (*jwt.Token, error) {
		return &jwt.Token{Valid: false}, nil
	}

	req, err := http.NewRequest("GET", "/?app-token=test-token", nil)
	assert.NoError(t, err)

	DashboardEntriesTree(rr, req)
	assert.Equal(t, http.StatusForbidden, rr.Code)
}

func TestDashboardEntriesTreeErrNoToken(t *testing.T) {
	rr, cleanup := setupDashboardEntriesTreeTest()
	defer cleanup()

	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	DashboardEntriesTree(rr, req)
	assert.Equal(t, http.StatusForbidden, rr.Code)
}

func TestDashboardEntriesTreeErrNoTemplate(t *testing.T) {
	rr, cleanup := setupDashboardEntriesTreeTest()
	defer cleanup()

	getFs = func() FS { return fstest.MapFS{} }

	req, err := http.NewRequest("GET", "/?app-token=test-token", nil)
	assert.NoError(t, err)

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

	req, err := http.NewRequest("GET", "/?app-token=test-token", nil)
	assert.NoError(t, err)

	DashboardEntriesTree(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "", rr.Body.String())
}

func TestDashboardEntriesTreeErrGetPublicKey(t *testing.T) {
	rr, cleanup := setupDashboardEntriesTreeTest()
	defer cleanup()

	csSdkGetUrl = func(useManagementToken bool) string {
		return "https://bogus\\"
	}

	req, err := http.NewRequest("GET", "/?app-token=test-token", nil)
	assert.NoError(t, err)

	DashboardEntriesTree(rr, req)
	assert.Equal(t, http.StatusForbidden, rr.Code)
}

func TestDashboardEntriesTreeErrJwtParseErr(t *testing.T) {
	rr, cleanup := setupDashboardEntriesTreeTest()
	defer cleanup()

	jwtParse = jwt.Parse

	req, err := http.NewRequest("GET", "/?app-token=bogus-token", nil)
	assert.NoError(t, err)

	DashboardEntriesTree(rr, req)
	assert.Equal(t, http.StatusForbidden, rr.Code)
}

func TestDashboardEntriesTreeErrJwtParse(t *testing.T) {
	rr, cleanup := setupDashboardEntriesTreeTest()
	defer cleanup()

	csSdkGetUrl = func(useManagementToken bool) string {
		return "bogus"
	}

	req, err := http.NewRequest("GET", "/?app-token=bogus-token", nil)
	assert.NoError(t, err)

	DashboardEntriesTree(rr, req)
	assert.Equal(t, http.StatusForbidden, rr.Code)
}

func TestDashboardEntriesTreeErrBody(t *testing.T) {
	rr, cleanup := setupDashboardEntriesTreeTest()
	defer cleanup()

	csSdkGetUrl = func(useManagementToken bool) string {
		return "nil-body"
	}

	req, err := http.NewRequest("GET", "/?app-token=bogus-token", nil)
	assert.NoError(t, err)

	DashboardEntriesTree(rr, req)
	assert.Equal(t, http.StatusForbidden, rr.Code)
}

func TestDashboardEntriesTreeErrJsonUnmarshal(t *testing.T) {
	rr, cleanup := setupDashboardEntriesTreeTest()
	defer cleanup()

	csSdkGetUrl = func(useManagementToken bool) string {
		return "json-err"
	}

	req, err := http.NewRequest("GET", "/?app-token=bogus-token", nil)
	assert.NoError(t, err)

	DashboardEntriesTree(rr, req)
	assert.Equal(t, http.StatusForbidden, rr.Code)
}

func TestDashboardEntriesTreeErrNoSigningKey(t *testing.T) {
	rr, cleanup := setupDashboardEntriesTreeTest()
	defer cleanup()

	csSdkGetUrl = func(useManagementToken bool) string {
		return "json-empty"
	}

	req, err := http.NewRequest("GET", "/?app-token=bogus-token", nil)
	assert.NoError(t, err)

	DashboardEntriesTree(rr, req)
	assert.Equal(t, http.StatusForbidden, rr.Code)
}
