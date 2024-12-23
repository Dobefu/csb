package api

import (
	"github.com/Dobefu/csb/cmd/cs_sdk"
)

func GetGlobalFields() (map[string]interface{}, error) {
	data, err := cs_sdk.Request("global_fields", "GET", nil, false)

	if err != nil {
		return nil, err
	}

	return data, nil
}
