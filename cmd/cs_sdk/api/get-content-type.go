package api

import (
	"fmt"

	"github.com/Dobefu/csb/cmd/cs_sdk"
)

func GetContentType(id string) interface{} {
	contentType, err := cs_sdk.Request(
		fmt.Sprintf("content_types/%s", id),
		"GET",
		nil,
	)

	if err != nil {
		return nil
	}

	return contentType
}
