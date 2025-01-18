package cs_sdk

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

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
		if req.URL.Path == "/v3/test-path" {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(`{"key":"value"}`)),
			}, nil
		}

		return &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       io.NopCloser(bytes.NewBufferString(`{}`)),
		}, errors.New("not found")
	},
}

func TestRequestRaw(t *testing.T) {
	originalClient := httpClient
	defer func() { httpClient = originalClient }()

	httpClient = mockClient

	response, err := RequestRaw("test-path", "GET", nil, false)

	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestRequestRawWithBody(t *testing.T) {
	originalClient := httpClient
	defer func() { httpClient = originalClient }()

	httpClient = mockClient

	response, err := RequestRaw("test-path", "GET", map[string]interface{}{}, false)

	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestRequestRawWithManagementToken(t *testing.T) {
	originalClient := httpClient
	defer func() { httpClient = originalClient }()

	httpClient = mockClient

	response, err := RequestRaw("test-path", "GET", nil, true)

	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestRequestRawErrInvalidBody(t *testing.T) {
	originalClient := httpClient
	defer func() { httpClient = originalClient }()

	httpClient = mockClient

	response, err := RequestRaw("test-path", "GET", map[string]interface{}{"invalid": func() { /* Invalid */ }}, false)

	assert.NotEqual(t, nil, err)
	assert.NotEqual(t, nil, response)
}

func TestRequestRawErrInvalidRequest(t *testing.T) {
	originalClient := httpClient
	defer func() { httpClient = originalClient }()

	httpClient = mockClient

	response, err := RequestRaw("test-path", "BOGUS@", nil, false)

	assert.NotEqual(t, nil, err)
	assert.NotEqual(t, nil, response)
}

func TestRequestRawErrNotFound(t *testing.T) {
	originalClient := httpClient
	defer func() { httpClient = originalClient }()

	httpClient = mockClient

	response, err := RequestRaw("bogus", "GET", nil, false)

	assert.NotEqual(t, nil, err)
	assert.NotEqual(t, nil, response)
}
