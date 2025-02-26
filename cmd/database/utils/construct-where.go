package utils

import (
	"fmt"
	"strings"

	"github.com/Dobefu/csb/cmd/database/structs"
	"github.com/Dobefu/csb/cmd/logger"
)

func ConstructWhere(where []structs.QueryWhere) (string, []any) {
	var whereStrings []string
	var whereArgs []any

	for _, whereSingle := range where {
		operator, err := ParseOperator(whereSingle.Operator)

		if err != nil {
			logger.Error(err.Error())
			continue
		}

		whereString := fmt.Sprintf(
			"%s %s %v",
			whereSingle.Name,
			operator,
			"?",
		)

		whereStrings = append(whereStrings, whereString)
		whereArgs = append(whereArgs, whereSingle.Value)
	}

	return fmt.Sprintf("WHERE %s", strings.Join(whereStrings, " AND ")), whereArgs
}
