package query

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/database/structs"
	"github.com/Dobefu/csb/cmd/database/utils"
)

func QueryRows(table string, fields []string, where []structs.QueryWhere) (*sql.Rows, error) {
	fieldsString := strings.Join(fields, ", ")

	switch os.Getenv("DB_TYPE") {
	case "mysql":
		return queryRowsMysql(table, fieldsString, where)
	case "sqlite3":
		return queryRowsSqlite3(table, fieldsString, where)
	default:
		return nil, nil
	}
}

func queryRowsMysql(table string, fields string, where []structs.QueryWhere) (*sql.Rows, error) {
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

	return database.DB.Query(strings.Join(sql, " "), args...)
}

func queryRowsSqlite3(table string, fields string, where []structs.QueryWhere) (*sql.Rows, error) {
	return queryRowsMysql(table, fields, where)
}
