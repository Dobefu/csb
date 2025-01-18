package check_health

import (
	"errors"
	"testing"

	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/Dobefu/csb/cmd/database"
	"github.com/stretchr/testify/assert"
)

var err error

func TestMain(t *testing.T) {
	databaseConnect = func() error { return nil }
	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		return map[string]interface{}{
			"label": map[string]interface{}{
				"uid": "dummy label"},
		}, nil
	}

	defer func() { databaseConnect = database.Connect }()
	defer func() { csSdkRequest = cs_sdk.Request }()

	err = Main()
	assert.Equal(t, nil, err)
}

func TestMainErrNoDatabase(t *testing.T) {
	databaseConnect = func() error { return errors.New("no database") }
	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		return map[string]interface{}{}, nil
	}

	defer func() { databaseConnect = database.Connect }()
	defer func() { csSdkRequest = cs_sdk.Request }()

	err = Main()
	assert.NotEqual(t, nil, err)
}

func TestMainErrNoLabel(t *testing.T) {
	databaseConnect = func() error { return nil }
	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		return map[string]interface{}{}, nil
	}

	defer func() { databaseConnect = database.Connect }()
	defer func() { csSdkRequest = cs_sdk.Request }()

	err = Main()
	assert.NotEqual(t, nil, err)
}

func TestCheckCsSdkErrLabelCreate(t *testing.T) {
	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		return nil, errors.New("cannot connect")
	}

	defer func() { csSdkRequest = cs_sdk.Request }()

	err = checkCsSdk()
	assert.NotEqual(t, nil, err)
}

func TestCheckCsSdkErrLabelDelete(t *testing.T) {
	csSdkRequest = func(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
		if method == "DELETE" {
			return nil, errors.New("cannot connect")
		}

		return map[string]interface{}{
			"label": map[string]interface{}{
				"uid": "dummy label"},
		}, nil
	}

	defer func() { csSdkRequest = cs_sdk.Request }()

	err = checkCsSdk()
	assert.NotEqual(t, nil, err)
}
