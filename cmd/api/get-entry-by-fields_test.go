package api

import (
	"testing"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/database/query"
	"github.com/Dobefu/csb/cmd/database/structs"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/Dobefu/csb/cmd/migrate_db"
	"github.com/stretchr/testify/assert"
)

func TestGetEntryByFields(t *testing.T) {
	init_env.Main("../../.env.test")
	err := database.Connect()
	assert.Equal(t, nil, err)

	where := []structs.QueryWhere{
		{
			Name:     "uid",
			Value:    "testing",
			Operator: db_structs.EQUALS,
		},
	}

	_, err = GetEntryByFields(where)
	assert.NotEqual(t, nil, err)

	err = migrate_db.Main(true)
	assert.Equal(t, nil, err)

	err = insertPage()
	assert.Equal(t, nil, err)

	route, err := GetEntryByFields(where)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, route)
}

func insertPage() error {
	return query.Insert("routes", []db_structs.QueryValue{
		{
			Name:  "id",
			Value: "testingen",
		},
		{
			Name:  "uid",
			Value: "testing",
		},
		{
			Name:  "content_type",
			Value: "basic_page",
		},
		{
			Name:  "locale",
			Value: "en",
		},
		{
			Name:  "slug",
			Value: "/testing",
		},
		{
			Name:  "url",
			Value: "/testing",
		},
		{
			Name:  "parent",
			Value: "parent_uid",
		},
	})
}
