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

func QueryRow(table string, fields []string, where []structs.QueryWhere) *sql.Row {
	var (
		sql  string
		args []any
	)

	fieldsString := strings.Join(fields, ", ")
	dbType := os.Getenv("DB_TYPE")

	switch dbType {
	case "mysql":
		sql, args = queryRowMysql(table, fieldsString, where)
	case "sqlite3":
		sql, args = queryRowSqlite3(table, fieldsString, where)
	case "postgres":
		sql, args = queryRowPostgres(table, fieldsString, where)
	default:
		logger.Fatal(
			"The database type %s has no corresponding QueryRow function",
			dbType,
		)

		return nil
	}

	return database.DB.QueryRow(sql, args...)
}

func queryRowMysql(table string, fields string, where []structs.QueryWhere) (string, []any) {
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

	sql = append(sql, "LIMIT 1")
	return strings.Join(sql, " "), args
}

func queryRowSqlite3(table string, fields string, where []structs.QueryWhere) (string, []any) {
	return queryRowMysql(table, fields, where)
}

func queryRowPostgres(table string, fields string, where []structs.QueryWhere) (string, []any) {
	sql, args := queryRowMysql(table, fields, where)

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
