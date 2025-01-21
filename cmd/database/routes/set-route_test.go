package routes

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/Dobefu/csb/cmd/cs_sdk/utils"
	"github.com/Dobefu/csb/cmd/database/query"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
	"github.com/stretchr/testify/assert"
)

func setupSetRouteTest() func() {
	queryUpsert = func(table string, values []db_structs.QueryValue) error {
		return nil
	}

	utilsGenerateId = func(uid string, locale string) string {
		return fmt.Sprintf("%s%s", uid, locale)
	}

	return func() {
		queryUpsert = query.Upsert
		utilsGenerateId = utils.GenerateId
	}
}

func TestSetRouteSuccess(t *testing.T) {
	cleanup := setupSetRouteTest()
	defer cleanup()

	err := SetRoute(structs.Route{})
	assert.NoError(t, err)
}

func TestSetRouteErrUpsert(t *testing.T) {
	cleanup := setupSetRouteTest()
	defer cleanup()

	queryUpsert = func(table string, values []db_structs.QueryValue) error {
		return errors.New("cannot upsert route")
	}

	err := SetRoute(structs.Route{})
	assert.EqualError(t, err, "cannot upsert route")
}
