package utils

import (
	"fmt"

	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
)

func GenerateId(route structs.Route) string {
	return fmt.Sprintf("%s%s", route.Uid, route.Locale)
}
