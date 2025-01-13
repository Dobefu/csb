package utils

import (
	"fmt"
)

func GenerateId(uid string, locale string) string {
	return fmt.Sprintf("%s%s", uid, locale)
}
