package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogMessage(t *testing.T) {
	SetLogLevel(LOG_FATAL)
	assert.NotEmpty(t, logMessage(255, "test"))
}

func TestLogVerbose(t *testing.T) {
	SetLogLevel(LOG_VERBOSE)
	assert.NotEmpty(t, Verbose("test"))
	SetLogLevel(LOG_INFO)
	assert.Empty(t, Verbose("test"))
	SetLogLevel(LOG_WARNING)
	assert.Empty(t, Verbose("test"))
	SetLogLevel(LOG_ERROR)
	assert.Empty(t, Verbose("test"))
	SetLogLevel(LOG_FATAL)
	assert.Empty(t, Verbose("test"))
}

func TestLogInfo(t *testing.T) {
	SetLogLevel(LOG_VERBOSE)
	assert.NotEmpty(t, Info("test"))
	SetLogLevel(LOG_INFO)
	assert.NotEmpty(t, Info("test"))
	SetLogLevel(LOG_WARNING)
	assert.Empty(t, Info("test"))
	SetLogLevel(LOG_ERROR)
	assert.Empty(t, Info("test"))
	SetLogLevel(LOG_FATAL)
	assert.Empty(t, Info("test"))
}

func TestLogWarning(t *testing.T) {
	SetLogLevel(LOG_VERBOSE)
	assert.NotEmpty(t, Warning("test"))
	SetLogLevel(LOG_INFO)
	assert.NotEmpty(t, Warning("test"))
	SetLogLevel(LOG_WARNING)
	assert.NotEmpty(t, Warning("test"))
	SetLogLevel(LOG_ERROR)
	assert.Empty(t, Warning("test"))
	SetLogLevel(LOG_FATAL)
	assert.Empty(t, Warning("test"))
}

func TestLogError(t *testing.T) {
	SetLogLevel(LOG_VERBOSE)
	assert.NotEmpty(t, Error("test"))
	SetLogLevel(LOG_INFO)
	assert.NotEmpty(t, Error("test"))
	SetLogLevel(LOG_WARNING)
	assert.NotEmpty(t, Error("test"))
	SetLogLevel(LOG_ERROR)
	assert.NotEmpty(t, Error("test"))
	SetLogLevel(LOG_FATAL)
	assert.Empty(t, Error("test"))
}

func TestLogFatal(t *testing.T) {
	SetLogLevel(LOG_VERBOSE)
	assert.NotEmpty(t, Fatal("test"))
	SetLogLevel(LOG_INFO)
	assert.NotEmpty(t, Fatal("test"))
	SetLogLevel(LOG_WARNING)
	assert.NotEmpty(t, Fatal("test"))
	SetLogLevel(LOG_ERROR)
	assert.NotEmpty(t, Fatal("test"))
	SetLogLevel(LOG_FATAL)
	assert.NotEmpty(t, Fatal("test"))
}
