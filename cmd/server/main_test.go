package server

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	listenAndServe = func(addr string, handler http.Handler) error {
		return nil
	}

	defer func() { listenAndServe = http.ListenAndServe }()

	err := Start(40000)
	assert.Equal(t, nil, err)
}

func TestStartErr(t *testing.T) {
	listenAndServe = func(addr string, handler http.Handler) error {
		return errors.New("")
	}

	defer func() { listenAndServe = http.ListenAndServe }()

	err := Start(40000)
	assert.NotEqual(t, nil, err)
}
