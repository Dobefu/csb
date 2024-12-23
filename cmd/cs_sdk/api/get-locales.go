package api

import (
	"github.com/Dobefu/csb/cmd/cs_sdk"
)

func GetLocales() interface{} {
	locales, err := cs_sdk.Request(
		"locales",
		"GET",
		nil,
		true,
	)

	if err != nil {
		return nil
	}

	return locales
}
