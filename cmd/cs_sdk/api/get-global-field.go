package api

import (
	"fmt"

	"github.com/Dobefu/csb/cmd/cs_sdk"
)

func GetGlobalField(id string) interface{} {
	globalField, err := cs_sdk.Request(
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
