package utils

import (
	"errors"
	"fmt"
	"os"

	db_structs "github.com/Dobefu/csb/cmd/database/structs"
	"github.com/Dobefu/csb/cmd/logger"
)

func ParseOperator(op db_structs.Operator) (string, error) {
	dbType := os.Getenv("DB_TYPE")

	switch dbType {
	case "mysql":
		return parseOperatorMysql(op)
	case "sqlite3":
		return parseOperatorSqlite3(op)
	case "postgres":
		return parseOperatorPostgres(op)
	default:
		msg := fmt.Sprintf(
			"The database type %s has no corresponding QueryRows function",
			dbType,
		)

		logger.Fatal(msg)
		return "", errors.New(msg)
	}
}

func parseOperatorMysql(op db_structs.Operator) (string, error) {
	switch op {
	case db_structs.EQUALS:
		return "=", nil
	case db_structs.NOT_EQUALS:
		return "<>", nil
	default:
		return "", errors.New("the used operator is unsupported")
	}
}

func parseOperatorSqlite3(op db_structs.Operator) (string, error) {
	return parseOperatorMysql(op)
}

func parseOperatorPostgres(op db_structs.Operator) (string, error) {
	return parseOperatorMysql(op)
}
