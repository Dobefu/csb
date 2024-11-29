package utils

import (
	"fmt"
	"strings"

	"github.com/Dobefu/csb/cmd/database/structs"
)

func ConstructWhere(where []structs.QueryWhere) (string, []any) {
	var whereStrings []string
	var whereArgs []any

	for _, whereSingle := range where {
		whereStrings = append(whereStrings, fmt.Sprintf("%s = %v", whereSingle.Name, "?"))
		whereArgs = append(whereArgs, whereSingle.Value)
	}

	return fmt.Sprintf("WHERE %s", strings.Join(whereStrings, " AND ")), whereArgs
}
