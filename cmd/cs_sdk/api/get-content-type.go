package api

import (
	"fmt"
)

func GetContentType(id string) interface{} {
	contentType, err := csSdkRequest(
		fmt.Sprintf("content_types/%s", id),
		"GET",
		nil,
		false,
	)

	if err != nil {
		return nil
	}

	return contentType
}
