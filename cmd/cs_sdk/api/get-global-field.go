package api

import (
	"fmt"
)

func GetGlobalField(id string) interface{} {
	globalField, err := csSdkRequest(
		fmt.Sprintf("global_fields/%s", id),
		"GET",
		nil,
		false,
	)

	if err != nil {
		return nil
	}

	return globalField
}
