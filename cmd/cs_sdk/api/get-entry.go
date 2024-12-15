package api

import (
	"fmt"

	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
)

func GetEntry(route structs.Route) (interface{}, error) {
	path := fmt.Sprintf(
		"content_types/%s/entries/%s?locale=%s",
		route.ContentType,
		route.Uid,
		route.Locale,
	)

	res, err := cs_sdk.Request(path, "GET", nil)

	if err != nil {
		return nil, err
	}

	return res["entry"], nil
}
