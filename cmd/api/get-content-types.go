package api

import (
	"github.com/Dobefu/csb/cmd/cs_sdk"
)

func GetContentTypes() (map[string]interface{}, error) {
	data, err := cs_sdk.Request("content_types", "GET", nil)

	if err != nil {
		return nil, err
	}

	return data, nil
}
