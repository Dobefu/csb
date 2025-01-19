package cli

import (
	"errors"
	"os"
	"testing"

	"github.com/Dobefu/csb/cmd/cs_sdk/api"
	"github.com/stretchr/testify/assert"
)

func setupTestCreateContentType(t *testing.T) (*os.File, func()) {
	tmpfile, err := os.CreateTemp("", "TestCreateContentType")
	assert.Equal(t, nil, err)

	apiCreateContentType = func(name string, id string, withFields bool) error {
		return nil
	}

	oldStdin := os.Stdin
	os.Stdin = tmpfile

	cleanup := func() {
		apiCreateContentType = api.CreateContentType
		os.Remove(tmpfile.Name())
		os.Stdin = oldStdin
	}

	return tmpfile, cleanup
}

func TestCreateContentTypeSuccessDryRun(t *testing.T) {
	err := CreateContentType(true, "Content Type", "content_type", true)
	assert.NoError(t, err)
}

func TestCreateContentTypeSuccess(t *testing.T) {
	_, cleanup := setupTestCreateContentType(t)
	defer cleanup()

	err := CreateContentType(false, "Content Type", "content_type", true)
	assert.NoError(t, err)
}

func TestCreateContentTypeErrCreate(t *testing.T) {
	_, cleanup := setupTestCreateContentType(t)
	defer cleanup()

	apiCreateContentType = func(name string, id string, withFields bool) error {
		return errors.New("cannot create content type")
	}

	err := CreateContentType(false, "Content Type", "content_type", true)
	assert.Error(t, err)
}

func TestCreateContentTypeErrNoName(t *testing.T) {
	err := CreateContentType(true, "", "content_type", true)
	assert.Error(t, err)
}

func TestCreateContentTypeErrNoMachineName(t *testing.T) {
	err := CreateContentType(true, "Content Type", "", true)
	assert.Error(t, err)
}

func TestCreateContentTypeSuccessNameConfirm(t *testing.T) {
	tmpfile, cleanup := setupTestCreateContentType(t)
	defer cleanup()

	_, err := tmpfile.WriteAt([]byte("Content Type\n"), 0)
	assert.NoError(t, err)

	_, err = tmpfile.Seek(0, 0)
	assert.NoError(t, err)

	err = CreateContentType(true, "", "content_type", true)
	assert.NoError(t, err)
}

func TestCreateContentTypeSuccessMachineNameConfirm(t *testing.T) {
	tmpfile, cleanup := setupTestCreateContentType(t)
	defer cleanup()

	_, err := tmpfile.WriteAt([]byte("content_type\n"), 0)
	assert.Equal(t, nil, err)

	_, err = tmpfile.Seek(0, 0)
	assert.Equal(t, nil, err)

	err = CreateContentType(true, "Content Type", "", true)
	assert.Equal(t, nil, err)
}

func TestCreateContentTypeErrNameEmpty(t *testing.T) {
	tmpfile, cleanup := setupTestCreateContentType(t)
	defer cleanup()

	_, err := tmpfile.WriteAt([]byte("\n"), 0)
	assert.Equal(t, nil, err)

	_, err = tmpfile.Seek(0, 0)
	assert.Equal(t, nil, err)

	err = CreateContentType(true, "Content Type", "", false)
	assert.NotEqual(t, nil, err)
}

func TestCreateContentTypeErrMachineNameEmpty(t *testing.T) {
	tmpfile, cleanup := setupTestCreateContentType(t)
	defer cleanup()

	_, err := tmpfile.Seek(0, 0)
	assert.Equal(t, nil, err)

	err = CreateContentType(true, "", "content_type", true)
	assert.NotEqual(t, nil, err)
}
