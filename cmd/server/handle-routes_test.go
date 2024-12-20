package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/Dobefu/csb/cmd/migrate_db"
	"github.com/Dobefu/csb/cmd/remote_sync"
	"github.com/stretchr/testify/assert"
)

func TestHandleRoutes(t *testing.T) {
	init_env.Main("../../.env.test")

	var err error

	err = database.Connect()
	assert.Equal(t, nil, err)

	err = migrate_db.Main(false)
	assert.Equal(t, nil, err)

	err = remote_sync.Sync(false)
	assert.Equal(t, nil, err)

	mux := http.NewServeMux()
	HandleRoutes(mux, "")

	server := httptest.NewServer(mux)

	defer server.Close()

	var body map[string]interface{}

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "/"),
	)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, body["data"])
	assert.Equal(t, nil, body["error"])

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "/get-entry-by-url"),
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, body["data"])
	assert.Contains(t, body["error"], "missing required query params")

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "/get-entry-by-url?url=/&locale=en"),
	)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, body["data"])
	assert.Equal(t, nil, body["error"])

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "/get-entry-by-url?url=/bogus&locale=en"),
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, body["data"])
	assert.NotEqual(t, nil, body["error"])

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "/get-entry-by-uid"),
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, body["data"])
	assert.Contains(t, body["error"], "missing required query params")

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "/get-entry-by-uid?uid=blt0617c28651fb44bf&locale=en"),
	)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, body["data"])
	assert.Equal(t, nil, body["error"])

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "/get-entry-by-uid?uid=/bogus&locale=en"),
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, body["data"])
	assert.NotEqual(t, nil, body["error"])

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "/content-types"),
	)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, body["data"])
	assert.Equal(t, nil, body["error"])
}

func request(method string, path string) (body map[string]interface{}, err error) {
	req, err := http.NewRequest(
		method,
		path,
		nil,
	)

	if err != nil {
		return nil, err
	}

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(respBody, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}
