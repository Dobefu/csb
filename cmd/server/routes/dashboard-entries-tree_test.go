package routes

import (
	"bytes"
	"crypto/rsa"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/Dobefu/csb/cmd/database/query"
	"github.com/Dobefu/csb/cmd/database/structs"
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

func setupDashboardEntriesTreeTest(t *testing.T) (sqlmock.Sqlmock, *httptest.ResponseRecorder, func()) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

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

	queryQueryRows = func(table string, fields []string, where []structs.QueryWhere) (*sql.Rows, error) {
		return db.Query("SELECT (.+) FROM routes", nil)
	}

	return mock, rr, func() {
		getFs = func() FS { return content }
		httpClient = originalClient
		csSdkGetUrl = cs_sdk.GetUrl
		jwtParse = jwt.Parse
		jwtParseRSAPublicKeyFromPEM = jwt.ParseRSAPublicKeyFromPEM
		queryQueryRows = query.QueryRows
	}
}

func TestDashboardEntriesTreeSuccess(t *testing.T) {
	mock, rr, cleanup := setupDashboardEntriesTreeTest(t)
	defer cleanup()

	jwtParse = func(tokenString string, keyFunc jwt.Keyfunc, options ...jwt.ParserOption) (*jwt.Token, error) {
		key := &jwt.Token{Valid: true, Claims: jwt.MapClaims{"stack_api_key": ""}}
		_, err := keyFunc(key)
		assert.NoError(t, err)

		return key, nil
	}

	cols := []string{"uid", "parent", "title"}
	mock.ExpectQuery("SELECT (.+) FROM routes").WillReturnRows(
		sqlmock.NewRows(cols).AddRow("uid", "parent", "title"),
	)

	req, err := http.NewRequest("GET", "/?app-token=some.test.token", nil)
	assert.NoError(t, err)

	DashboardEntriesTree(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.NotEmpty(t, rr.Body.String())
}

func TestDashboardEntriesTreeErrNoClaim(t *testing.T) {
	mock, rr, cleanup := setupDashboardEntriesTreeTest(t)
	defer cleanup()

	jwtParse = func(tokenString string, keyFunc jwt.Keyfunc, options ...jwt.ParserOption) (*jwt.Token, error) {
		key := &jwt.Token{Valid: true}
		_, err := keyFunc(key)
		assert.NoError(t, err)

		return key, nil
	}

	cols := []string{"uid", "parent", "title"}
	mock.ExpectQuery("SELECT (.+) FROM routes").WillReturnRows(
		sqlmock.NewRows(cols).AddRow("uid", "parent", "title"),
	)

	req, err := http.NewRequest("GET", "/?app-token=some.test.token", nil)
	assert.NoError(t, err)

	DashboardEntriesTree(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestDashboardEntriesTreeErrParsePublicKeyFromPem(t *testing.T) {
	_, rr, cleanup := setupDashboardEntriesTreeTest(t)
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
	_, rr, cleanup := setupDashboardEntriesTreeTest(t)
	defer cleanup()

	jwtParse = func(tokenString string, keyFunc jwt.Keyfunc, options ...jwt.ParserOption) (*jwt.Token, error) {
		return &jwt.Token{Valid: false}, nil
	}

	req, err := http.NewRequest("GET", "/?app-token=test-token", nil)
	assert.NoError(t, err)

	DashboardEntriesTree(rr, req)
	assert.Equal(t, http.StatusForbidden, rr.Code)
}

func TestDashboardEntriesTreeErrGetData(t *testing.T) {
	_, rr, cleanup := setupDashboardEntriesTreeTest(t)
	defer cleanup()

	jwtParse = func(tokenString string, keyFunc jwt.Keyfunc, options ...jwt.ParserOption) (*jwt.Token, error) {
		key := &jwt.Token{Valid: true, Claims: jwt.MapClaims{"stack_api_key": ""}}
		_, err := keyFunc(key)
		assert.NoError(t, err)

		return key, nil
	}

	req, err := http.NewRequest("GET", "/?app-token=test-token", nil)
	assert.NoError(t, err)

	DashboardEntriesTree(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestDashboardEntriesTreeErrNoToken(t *testing.T) {
	_, rr, cleanup := setupDashboardEntriesTreeTest(t)
	defer cleanup()

	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	DashboardEntriesTree(rr, req)
	assert.Equal(t, http.StatusForbidden, rr.Code)
}

func TestDashboardEntriesTreeErrNoTemplate(t *testing.T) {
	mock, rr, cleanup := setupDashboardEntriesTreeTest(t)
	defer cleanup()

	jwtParse = func(tokenString string, keyFunc jwt.Keyfunc, options ...jwt.ParserOption) (*jwt.Token, error) {
		key := &jwt.Token{Valid: true, Claims: jwt.MapClaims{"stack_api_key": ""}}
		_, err := keyFunc(key)
		assert.NoError(t, err)

		return key, nil
	}

	getFs = func() FS { return fstest.MapFS{} }

	cols := []string{"uid", "parent", "title"}
	mock.ExpectQuery("SELECT (.+) FROM routes").WillReturnRows(
		sqlmock.NewRows(cols).AddRow("uid", "parent", "title"),
	)

	req, err := http.NewRequest("GET", "/?app-token=test-token", nil)
	assert.NoError(t, err)

	DashboardEntriesTree(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "", rr.Body.String())
}

func TestDashboardEntriesTreeErrInvalidTemplate(t *testing.T) {
	mock, rr, cleanup := setupDashboardEntriesTreeTest(t)
	defer cleanup()

	jwtParse = func(tokenString string, keyFunc jwt.Keyfunc, options ...jwt.ParserOption) (*jwt.Token, error) {
		key := &jwt.Token{Valid: true, Claims: jwt.MapClaims{"stack_api_key": ""}}
		_, err := keyFunc(key)
		assert.NoError(t, err)

		return key, nil
	}

	getFs = func() FS {
		return fstest.MapFS{
			"templates/dashboard-entries-tree.html.tmpl": {
				Data: []byte("{{ .Bogus }}"),
			},
		}
	}

	cols := []string{"uid", "parent", "title"}
	mock.ExpectQuery("SELECT (.+) FROM routes").WillReturnRows(
		sqlmock.NewRows(cols).AddRow("uid", "parent", "title"),
	)

	req, err := http.NewRequest("GET", "/?app-token=test-token", nil)
	assert.NoError(t, err)

	DashboardEntriesTree(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "", rr.Body.String())
}

func TestDashboardEntriesTreeErrGetPublicKey(t *testing.T) {
	_, rr, cleanup := setupDashboardEntriesTreeTest(t)
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
	_, rr, cleanup := setupDashboardEntriesTreeTest(t)
	defer cleanup()

	jwtParse = jwt.Parse

	req, err := http.NewRequest("GET", "/?app-token=bogus-token", nil)
	assert.NoError(t, err)

	DashboardEntriesTree(rr, req)
	assert.Equal(t, http.StatusForbidden, rr.Code)
}

func TestDashboardEntriesTreeErrJwtParse(t *testing.T) {
	_, rr, cleanup := setupDashboardEntriesTreeTest(t)
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
	_, rr, cleanup := setupDashboardEntriesTreeTest(t)
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
	_, rr, cleanup := setupDashboardEntriesTreeTest(t)
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
	_, rr, cleanup := setupDashboardEntriesTreeTest(t)
	defer cleanup()

	csSdkGetUrl = func(useManagementToken bool) string {
		return "json-empty"
	}

	req, err := http.NewRequest("GET", "/?app-token=bogus-token", nil)
	assert.NoError(t, err)

	DashboardEntriesTree(rr, req)
	assert.Equal(t, http.StatusForbidden, rr.Code)
}

func TestGetNestedEntriesErrNoRows(t *testing.T) {
	_, _, cleanup := setupDashboardEntriesTreeTest(t)
	defer cleanup()

	nestedEntries := getNestedEntries(map[string]interface{}{
		"test-uid": map[string]interface{}{
			"uid":    "test-uid",
			"title":  "test-title",
			"parent": "",
		},
		"test-uid2": map[string]interface{}{
			"uid":    "test-uid2",
			"title":  "test-title2",
			"parent": "test-uid",
		},
	})

	assert.Equal(t, map[string]interface{}(
		map[string]interface{}{
			"test-uid": map[string]interface{}{
				"children": []interface{}{
					map[string]interface{}{
						"children": []interface{}{},
						"parent":   "test-uid",
						"title":    "test-title2",
						"uid":      "test-uid2",
					},
				},
				"parent": "",
				"title":  "test-title",
				"uid":    "test-uid",
			},
		},
	), nestedEntries)
}

func TestGetEntriesErrNoRows(t *testing.T) {
	mock, _, cleanup := setupDashboardEntriesTreeTest(t)
	defer cleanup()

	cols := []string{"uid", "parent", "title"}
	mock.ExpectQuery("SELECT (.+) FROM routes").WillReturnRows(
		sqlmock.NewRows(cols).AddRow("uid", "parent", nil),
	)

	entries, err := getEntries()
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{}, entries)
}
