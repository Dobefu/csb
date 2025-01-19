package utils

import (
	"errors"
	"testing"

	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
	"github.com/stretchr/testify/assert"
)

type MockRows struct {
	nextCalls int
	scanError error
}

func (m *MockRows) Next() bool {
	m.nextCalls++
	return m.nextCalls == 1
}

func (m *MockRows) Scan(dest ...interface{}) error {
	if m.scanError != nil {
		return m.scanError
	}

	uid := dest[0].(*string)
	contentType := dest[1].(*string)
	locale := dest[2].(*string)
	slug := dest[3].(*string)
	url := dest[4].(*string)

	*uid = "test-uid"
	*contentType = "test-content-type"
	*locale = "test-locale"
	*slug = "test-slug"
	*url = "test-url"

	return nil
}

var mockQueryRows func(table string, columns []string, where []db_structs.QueryWhere) (queryRows, error)

func TestGetAltLocales(t *testing.T) {
	queryQueryRows = func(table string, columns []string, where []db_structs.QueryWhere) (queryRows, error) {
		return mockQueryRows(table, columns, where)
	}

	entry := structs.Route{
		Uid:    "test-uid",
		Locale: "en",
	}

	t.Run("successful retrieval", func(t *testing.T) {
		mockQueryRows = func(table string, columns []string, where []db_structs.QueryWhere) (queryRows, error) {
			return &MockRows{}, nil
		}

		results, err := GetAltLocales(entry, false)
		assert.NoError(t, err)
		assert.Len(t, results, 1)
	})

	t.Run("query error", func(t *testing.T) {
		mockQueryRows = func(table string, columns []string, where []db_structs.QueryWhere) (queryRows, error) {
			return nil, errors.New("query error")
		}

		results, err := GetAltLocales(entry, false)
		assert.Error(t, err)
		assert.Nil(t, results)
	})

	t.Run("scan error", func(t *testing.T) {
		mockQueryRows = func(_ string, _ []string, _ []db_structs.QueryWhere) (queryRows, error) {
			return &MockRows{scanError: errors.New("scan error")}, nil
		}

		results, err := GetAltLocales(entry, false)
		assert.NoError(t, err)
		assert.Len(t, results, 0)
	})
}
