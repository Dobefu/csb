package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/Dobefu/csb/cmd/migrate_db"
	"github.com/Dobefu/csb/cmd/remote_sync"
	"github.com/stretchr/testify/assert"
)

func TestHandleRoutes(t *testing.T) {
	var err error

	init_env.Main("../../.env.test")

	err = database.Connect()
	assert.Equal(t, nil, err)

	err = migrate_db.Main(true)
	assert.Equal(t, nil, err)

	err = remote_sync.Sync(true)
	assert.Equal(t, nil, err)

	mux := http.NewServeMux()
	HandleRoutes(mux, "")

	server := httptest.NewServer(mux)

	defer server.Close()

	var body map[string]interface{}

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, ""),
		true,
	)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, body["data"])
	assert.Equal(t, nil, body["error"])

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "bogus"),
		true,
	)
	assert.NotEqual(t, nil, err)
	assert.NotEqual(t, nil, body)

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "get-entry-by-url"),
		true,
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, body["data"])
	assert.Contains(t, body["error"], "missing required query params")

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "get-entry-by-url?url=/&locale=en"),
		true,
	)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, body["data"])
	assert.Equal(t, nil, body["error"])

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "get-entry-by-url?url=/bogus&locale=en"),
		true,
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, body["data"])
	assert.NotEqual(t, nil, body["error"])

	oldApiKey := os.Getenv("CS_API_KEY")
	os.Setenv("CS_API_KEY", "bogus")

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "get-entry-by-url?url=/&locale=en"),
		true,
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, body["data"])
	assert.NotEqual(t, nil, body["error"])

	os.Setenv("CS_API_KEY", oldApiKey)

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "get-entry-by-uid"),
		true,
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, body["data"])
	assert.Contains(t, body["error"], "missing required query params")

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "get-entry-by-uid?uid=blt0617c28651fb44bf&locale=en"),
		true,
	)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, body["data"])
	assert.Equal(t, nil, body["error"])

	oldApiKey = os.Getenv("CS_API_KEY")
	os.Setenv("CS_API_KEY", "bogus")

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "get-entry-by-uid?uid=blt0617c28651fb44bf&locale=en"),
		true,
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, body["data"])
	assert.NotEqual(t, nil, body["error"])

	os.Setenv("CS_API_KEY", oldApiKey)

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "get-entry-by-uid?uid=/bogus&locale=en"),
		true,
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, body["data"])
	assert.NotEqual(t, nil, body["error"])

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "content-types"),
		true,
	)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, body["data"])
	assert.Equal(t, nil, body["error"])

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "content-types"),
		false,
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, body["data"])
	assert.NotEqual(t, nil, body["error"])

	os.Setenv("DEBUG_AUTH_BYPASS", "1")

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "content-types"),
		false,
	)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, body["data"])
	assert.Equal(t, nil, body["error"])

	os.Setenv("DEBUG_AUTH_BYPASS", "")

	oldApiKey = os.Getenv("CS_API_KEY")
	os.Setenv("CS_API_KEY", "bogus")

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "content-types"),
		true,
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, body["data"])
	assert.NotEqual(t, nil, body["error"])

	os.Setenv("CS_API_KEY", oldApiKey)

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "content-type"),
		true,
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, body["data"])
	assert.Contains(t, body["error"], "missing required query params")

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "content-type?content_type=basic_page"),
		true,
	)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, body["data"])
	assert.Equal(t, nil, body["error"])

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "content-type?content_type=bogus"),
		true,
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, body["data"])
	assert.NotEqual(t, nil, body["error"])

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "global-fields"),
		true,
	)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, body["data"])
	assert.Equal(t, nil, body["error"])

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "global-fields"),
		false,
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, body["data"])
	assert.NotEqual(t, nil, body["error"])

	os.Setenv("DEBUG_AUTH_BYPASS", "1")

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "global-fields"),
		false,
	)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, body["data"])
	assert.Equal(t, nil, body["error"])

	os.Setenv("DEBUG_AUTH_BYPASS", "")

	oldApiKey = os.Getenv("CS_API_KEY")
	os.Setenv("CS_API_KEY", "bogus")

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "global-fields"),
		true,
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, body["data"])
	assert.NotEqual(t, nil, body["error"])

	os.Setenv("CS_API_KEY", oldApiKey)

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "locales"),
		true,
	)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, body["data"])
	assert.Equal(t, nil, body["error"])

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "locales"),
		false,
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, body["data"])
	assert.NotEqual(t, nil, body["error"])

	os.Setenv("DEBUG_AUTH_BYPASS", "1")

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "locales"),
		false,
	)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, body["data"])
	assert.Equal(t, nil, body["error"])

	os.Setenv("DEBUG_AUTH_BYPASS", "")

	oldApiKey = os.Getenv("CS_API_KEY")
	os.Setenv("CS_API_KEY", "bogus")

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "locales"),
		true,
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, body["data"])
	assert.NotEqual(t, nil, body["error"])

	os.Setenv("CS_API_KEY", oldApiKey)

	body, err = request(
		"POST",
		fmt.Sprintf("%s/%s", server.URL, "sync"),
		true,
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, body["error"])

	oldApiKey = os.Getenv("CS_API_KEY")
	os.Setenv("CS_API_KEY", "bogus")

	body, err = request(
		"POST",
		fmt.Sprintf("%s/%s", server.URL, "sync"),
		true,
	)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, body["error"])

	os.Setenv("CS_API_KEY", oldApiKey)

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "sitemap-data"),
		true,
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, body["error"])

	oldApiKey = os.Getenv("CS_API_KEY")
	os.Setenv("CS_API_KEY", "bogus")

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "sitemap-data"),
		true,
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, body["error"])

	os.Setenv("CS_API_KEY", oldApiKey)

	oldDb := os.Getenv("DB_CONN")
	os.Setenv("DB_CONN", "file:/")
	err = database.Connect()
	assert.Equal(t, nil, err)

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "sitemap-data"),
		true,
	)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, body["error"])

	os.Setenv("DB_CONN", oldDb)
	err = database.Connect()
	assert.Equal(t, nil, err)

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "translations"),
		true,
	)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, body["error"])

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "translations?locale=en"),
		true,
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, body["error"])

	oldApiKey = os.Getenv("CS_API_KEY")
	os.Setenv("CS_API_KEY", "bogus")

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "translations?locale=en"),
		true,
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, body["error"])

	os.Setenv("CS_API_KEY", oldApiKey)

	oldDb = os.Getenv("DB_CONN")
	os.Setenv("DB_CONN", "file:/")
	err = database.Connect()
	assert.Equal(t, nil, err)

	body, err = request(
		"GET",
		fmt.Sprintf("%s/%s", server.URL, "translations?locale=en"),
		true,
	)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, body["error"])

	os.Setenv("DB_CONN", oldDb)
	err = database.Connect()
	assert.Equal(t, nil, err)
}

func request(method string, path string, withAuthToken bool) (body map[string]interface{}, err error) {
	req, err := http.NewRequest(
		method,
		path,
		nil,
	)

	if err != nil {
		return nil, err
	}

	if withAuthToken {
		req.Header = http.Header{
			"Authorization": {os.Getenv("CS_DELIVERY_TOKEN")},
		}
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
