package cs_sdk

import (
	"bytes"
	"errors"
	"fmt"
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

type BrokenReader struct{}

func (br *BrokenReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("failed reading")
}

func (br *BrokenReader) Close() error {
	return fmt.Errorf("failed closing")
}

var mockClient = &MockClient{
	DoFunc: func(req *http.Request) (*http.Response, error) {
		if req.URL.Path == "/v3/test-path" {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(`{"key":"value"}`)),
			}, nil
		}

		if req.Method == "PUT" {
			return &http.Response{
				StatusCode: http.StatusNotFound,
				Body:       io.NopCloser(bytes.NewBufferString(`{}`)),
			}, nil
		}

		if req.URL.Path == "/v3/read-body" {
			reader := BrokenReader{}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       &reader,
			}, nil
		}

		if req.URL.Path == "/v3/wrong-body" {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(`{"key":"value"`)),
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

func TestRequest(t *testing.T) {
	originalClient := httpClient
	defer func() { httpClient = originalClient }()

	httpClient = mockClient

	response, err := Request("test-path", "GET", nil, false)

	assert.Equal(t, nil, err)
	assert.Equal(t, map[string]interface{}{"key": "value"}, response)
}

func TestRequestErrNotFound(t *testing.T) {
	originalClient := httpClient
	defer func() { httpClient = originalClient }()

	httpClient = mockClient

	response, err := Request("bogus", "GET", nil, false)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, map[string]interface{}(nil), response)
}

func TestRequestErrNotOk(t *testing.T) {
	originalClient := httpClient
	defer func() { httpClient = originalClient }()

	httpClient = mockClient

	response, err := Request("bogus", "PUT", nil, false)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, map[string]interface{}(nil), response)
}

func TestRequestErrReadBody(t *testing.T) {
	originalClient := httpClient
	defer func() { httpClient = originalClient }()

	httpClient = mockClient

	response, err := Request("read-body", "GET", nil, false)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, map[string]interface{}(nil), response)
}

func TestRequestErrWrongBody(t *testing.T) {
	originalClient := httpClient
	defer func() { httpClient = originalClient }()

	httpClient = mockClient

	response, err := Request("wrong-body", "GET", nil, false)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, map[string]interface{}(nil), response)
}
