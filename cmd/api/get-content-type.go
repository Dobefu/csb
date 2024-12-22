package api

import (
	"fmt"

	"github.com/Dobefu/csb/cmd/cs_sdk"
)

func GetContentType(contentType string) (map[string]interface{}, error) {
	data, err := cs_sdk.Request(
		fmt.Sprintf("content_types/%s", contentType),
		"GET",
		nil,
	)

	if err != nil {
		return nil, err
	}

	return data, nil
}
