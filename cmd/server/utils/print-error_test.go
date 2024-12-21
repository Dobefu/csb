package utils

import (
	"errors"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintError(t *testing.T) {
	var respBody []byte
	var err error

	w := httptest.NewRecorder()

	err = errors.New("Test error")
	PrintError(w, err, false)
	respBody, err = io.ReadAll(w.Body)
	assert.Equal(t, nil, err)
	assert.Contains(t, string(respBody), "Test error")

	err = errors.New("Test error")
	PrintError(w, err, true)
	respBody, err = io.ReadAll(w.Body)
	assert.Equal(t, nil, err)
	assert.Contains(t, string(respBody), "Test error")
}
