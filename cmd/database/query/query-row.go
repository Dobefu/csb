package query

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/Dobefu/csb/cmd/database"
)

func QueryRow(table string, fields []string) *sql.Row {
	fieldsString := strings.Join(fields, ", ")

	switch os.Getenv("DB_TYPE") {
	case "mysql":
		return queryRowMysql(table, fieldsString)
	case "sqlite3":
		return queryRowSqlite3(table, fieldsString)
	default:
		return nil
	}
}

func queryRowMysql(table string, fields string) *sql.Row {
	sql := fmt.Sprintf(
		"SELECT %s FROM %s LIMIT 1",
		fields,
		table,
	)

	return database.DB.QueryRow(sql)
}

func queryRowSqlite3(table string, fields string) *sql.Row {
	return queryRowMysql(table, fields)
}
