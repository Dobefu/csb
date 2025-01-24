package assets

import (
	"errors"
	"testing"

	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/Dobefu/csb/cmd/cs_sdk/utils"
	"github.com/Dobefu/csb/cmd/database/query"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
	"github.com/stretchr/testify/assert"
)

func setupSetAssetTest() func() {
	queryUpsert = func(table string, values []db_structs.QueryValue) error {
		return nil
	}

	return func() {
		queryUpsert = query.Upsert
		utilsGenerateId = utils.GenerateId
	}
}

func TestSetAssetSuccess(t *testing.T) {
	cleanup := setupSetAssetTest()
	defer cleanup()

	err := SetAsset(structs.Asset{})
	assert.NoError(t, err)
}

func TestSetAssetErrUpsert(t *testing.T) {
	cleanup := setupSetAssetTest()
	defer cleanup()

	queryUpsert = func(table string, values []db_structs.QueryValue) error {
		return errors.New("cannot upsert asset")
	}

	err := SetAsset(structs.Asset{})
	assert.EqualError(t, err, "cannot upsert asset")
}
