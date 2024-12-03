package query

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/database/structs"
	"github.com/Dobefu/csb/cmd/database/utils"
	"github.com/Dobefu/csb/cmd/logger"
)

func QueryRows(table string, fields []string, where []structs.QueryWhere) (*sql.Rows, error) {
	var (
		sql  string
		args []any
	)

	fieldsString := strings.Join(fields, ", ")
	dbType := os.Getenv("DB_TYPE")

	switch dbType {
	case "mysql":
		sql, args = queryRowsMysql(table, fieldsString, where)
	case "sqlite3":
		sql, args = queryRowsSqlite3(table, fieldsString, where)
	case "postgres":
		sql, args = queryRowsPostgres(table, fieldsString, where)
	default:
		logger.Fatal(
			"The database type %s has no corresponding QueryRows function",
			dbType,
		)

		return nil, nil
	}

	return database.DB.Query(sql, args...)
}

func queryRowsMysql(table string, fields string, where []structs.QueryWhere) (string, []any) {
	sql := []string{fmt.Sprintf(
		"SELECT %s FROM %s",
		fields,
		table,
	)}

	var args []any

	if where != nil {
		whereString, newArgs := utils.ConstructWhere(where)

		sql = append(sql, whereString)
		args = append(args, newArgs...)
	}

	return strings.Join(sql, " "), args
}

func queryRowsSqlite3(table string, fields string, where []structs.QueryWhere) (string, []any) {
	return queryRowsMysql(table, fields, where)
}

func queryRowsPostgres(table string, fields string, where []structs.QueryWhere) (string, []any) {
	sql, args := queryRowsMysql(table, fields, where)

	iteration := 0

	for {
		iteration += 1
		newSql := strings.Replace(sql, "?", fmt.Sprintf("$%d", iteration), 1)

		if newSql == sql {
			break
		}

		sql = newSql
	}

	return sql, args
}
