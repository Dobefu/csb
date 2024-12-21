package cli

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateContentType(t *testing.T) {
	var err error

	tmpfile, err := os.CreateTemp("", "TestCreateContentType")
	assert.Equal(t, nil, err)

	defer os.Remove(tmpfile.Name())

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	os.Stdin = tmpfile

	_, err = tmpfile.WriteAt([]byte("\n"), 0)
	assert.Equal(t, nil, err)

	_, err = tmpfile.Seek(0, 0)
	assert.Equal(t, nil, err)

	err = CreateContentType(true, "Content Type", "content_type")
	assert.Equal(t, nil, err)

	err = CreateContentType(true, "", "content_type")
	assert.NotEqual(t, nil, err)

	err = CreateContentType(true, "Content Type", "")
	assert.NotEqual(t, nil, err)

	_, err = tmpfile.WriteAt([]byte("Content Type\n"), 0)
	assert.Equal(t, nil, err)

	_, err = tmpfile.Seek(0, 0)
	assert.Equal(t, nil, err)

	err = CreateContentType(true, "", "content_type")
	assert.Equal(t, nil, err)

	_, err = tmpfile.WriteAt([]byte("content_type\n"), 0)
	assert.Equal(t, nil, err)

	_, err = tmpfile.Seek(0, 0)
	assert.Equal(t, nil, err)

	err = CreateContentType(true, "Content Type", "")
	assert.Equal(t, nil, err)
}
