package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Dobefu/csb/cmd/database/structs"
)

func ConstructWhere(where []structs.QueryWhere) string {
	var whereStrings []string

	for _, whereSingle := range where {
		val := whereSingle.Value

		if reflect.TypeOf(whereSingle.Value) == reflect.TypeOf("") {
			val = fmt.Sprintf("\"%s\"", val)
		}

		whereStrings = append(whereStrings, fmt.Sprintf("%s = %v", whereSingle.Name, val))
	}

	return fmt.Sprintf("WHERE %s", strings.Join(whereStrings, " AND "))
}
