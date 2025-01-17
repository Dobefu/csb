package utils

import (
	"errors"
	"testing"

	"github.com/Dobefu/csb/cmd/api"
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/stretchr/testify/assert"
)

func TestGetBreadcrumbs(t *testing.T) {
	apiGetEntryByUid = mockGetEntryByUid
	apiGetEntryByUrl = mockGetEntryByUrl

	defer func() {
		apiGetEntryByUid = api.GetEntryByUid
		apiGetEntryByUrl = api.GetEntryByUrl
	}()

	tests := []struct {
		name     string
		entry    structs.Route
		expected []structs.Route
	}{
		{
			name:  "No parent",
			entry: structs.Route{Uid: "entry1", Parent: "", Url: "/entry1", Locale: "en"},
			expected: []structs.Route{
				{Uid: "home", Parent: "", Url: "/", Locale: "en"},
				{Uid: "entry1", Parent: "", Url: "/entry1", Locale: "en"},
			},
		},
		{
			name:  "With parents",
			entry: structs.Route{Uid: "entry2", Parent: "parent1", Url: "/entry2", Locale: "en"},
			expected: []structs.Route{
				{Uid: "home", Parent: "", Url: "/", Locale: "en"},
				{Uid: "parent2", Parent: "", Url: "/parent2", Locale: "en"},
				{Uid: "parent1", Parent: "parent2", Url: "/parent1", Locale: "en"},
				{Uid: "entry2", Parent: "parent1", Url: "/entry2", Locale: "en"},
			},
		},
		{
			name:  "Parent not found",
			entry: structs.Route{Uid: "entry3", Parent: "nonexistent", Url: "/entry3", Locale: "en"},
			expected: []structs.Route{
				{Uid: "home", Parent: "", Url: "/", Locale: "en"},
				{Uid: "entry3", Parent: "nonexistent", Url: "/entry3", Locale: "en"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetBreadcrumbs(tt.entry)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func mockGetEntryByUid(uid, locale string, flag bool) (structs.Route, error) {
	if uid == "parent1" {
		return structs.Route{Uid: "parent1", Parent: "parent2", Url: "/parent1", Locale: locale}, nil
	}

	if uid == "parent2" {
		return structs.Route{Uid: "parent2", Parent: "", Url: "/parent2", Locale: locale}, nil
	}

	if uid == "nonexistent" {
		return structs.Route{}, errors.New("")
	}

	return structs.Route{}, errors.New("entry not found")
}

func mockGetEntryByUrl(url, locale string, flag bool) (structs.Route, error) {
	if url == "/" {
		return structs.Route{Uid: "home", Parent: "", Url: "/", Locale: locale}, nil
	}

	return structs.Route{}, errors.New("")
}
