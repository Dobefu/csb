package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/stretchr/testify/assert"
)

func TestAddResponseHeaders(t *testing.T) {
	var err error

	init_env.Main("../../../.env.test")

	mux := http.NewServeMux()
	AddResponseHeaders(mux)

	server := httptest.NewServer(mux)
	assert.NotEqual(t, nil, server)
	defer server.Close()

	mux.Handle(
		"/",
		AddResponseHeaders(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "")
		})),
	)

	err = request(
		fmt.Sprintf("%s/%s", server.URL, "/"),
		false,
	)
	assert.Equal(t, nil, err)
}
