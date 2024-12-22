package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/stretchr/testify/assert"
)

func TestRequireDeliveryToken(t *testing.T) {
	var err error

	init_env.Main("../../../.env.test")

	mux := http.NewServeMux()
	RequireDeliveryToken(mux)

	server := httptest.NewServer(mux)
	assert.NotEqual(t, nil, server)
	defer server.Close()

	mux.Handle(
		"/",
		RequireDeliveryToken(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "")
		})),
	)

	err = request(
		fmt.Sprintf("%s/%s", server.URL, "/"),
		false,
	)
	assert.Equal(t, nil, err)

	err = request(
		fmt.Sprintf("%s/%s", server.URL, "/"),
		true,
	)
	assert.Equal(t, nil, err)

	os.Setenv("DEBUG_AUTH_BYPASS", "1")

	err = request(
		fmt.Sprintf("%s/%s", server.URL, "/"),
		false,
	)
	assert.Equal(t, nil, err)

	os.Setenv("DEBUG_AUTH_BYPASS", "")

	oldAuthToken := os.Getenv("CS_DELIVERY_TOKEN")
	os.Setenv("CS_DELIVERY_TOKEN", "bogus")

	err = request(
		fmt.Sprintf("%s/%s", server.URL, "/"),
		false,
	)
	assert.Equal(t, nil, err)

	os.Setenv("CS_DELIVERY_TOKEN", "")

	err = request(
		fmt.Sprintf("%s/%s", server.URL, "/"),
		false,
	)
	assert.Equal(t, nil, err)

	os.Setenv("CS_DELIVERY_TOKEN", oldAuthToken)
}

func request(path string, withAuthToken bool) (err error) {
	req, err := http.NewRequest(
		"GET",
		path,
		nil,
	)

	if err != nil {
		return err
	}

	if withAuthToken {
		req.Header = http.Header{
			"Authorization": {os.Getenv("CS_DELIVERY_TOKEN")},
		}
	}

	client := http.Client{}
	_, err = client.Do(req)

	if err != nil {
		return err
	}

	return nil
}
