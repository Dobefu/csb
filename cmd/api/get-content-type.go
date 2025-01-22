package api

import (
	"fmt"

	"github.com/Dobefu/csb/cmd/cs_sdk"
)

var csSdkRequest = cs_sdk.Request

func GetContentType(contentType string) (map[string]interface{}, error) {
	data, err := csSdkRequest(
		fmt.Sprintf("content_types/%s", contentType),
		"GET",
		nil,
		false,
	)

	if err != nil {
		return nil, err
	}

	return data, nil
}
