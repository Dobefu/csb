package query

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/database/structs"
)

func Insert(table string, values []structs.QueryValue) error {
	var err error

	switch os.Getenv("DB_TYPE") {
	case "mysql":
		_, err = insertRowMysql(table, values)
	case "sqlite3":
		_, err = insertRowSqlite3(table, values)

	}

	return err
}

func insertRowMysql(table string, values []structs.QueryValue) (sql.Result, error) {
	sql := []string{fmt.Sprintf(
		"INSERT INTO %s",
		table,
	)}

	var valueNames []string
	var valuePlaceholders []string
	var args []any

	for _, value := range values {
		valueNames = append(valueNames, value.Name)
		valuePlaceholders = append(valuePlaceholders, "?")
		args = append(args, value.Value)
	}

	sql = append(sql, fmt.Sprintf("(%s)", strings.Join(valueNames, ", ")))
	sql = append(sql, fmt.Sprintf("VALUES (%s)", strings.Join(valuePlaceholders, ", ")))

	return database.DB.Exec(strings.Join(sql, " "), args...)
}

func insertRowSqlite3(table string, values []structs.QueryValue) (sql.Result, error) {
	return insertRowMysql(table, values)
}
