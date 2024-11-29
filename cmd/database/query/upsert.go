package query

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/database/structs"
)

func Upsert(table string, values []structs.QueryValue) error {
	var err error

	switch os.Getenv("DB_TYPE") {
	case "mysql":
		_, err = upsertRowMysql(table, values)
	case "sqlite3":
		_, err = upsertRowSqlite3(table, values)

	}

	return err
}

func upsertRowMysql(table string, values []structs.QueryValue) (sql.Result, error) {
	sql := []string{fmt.Sprintf(
		"INSERT INTO %s",
		table,
	)}

	var valueNames []string
	var valuePlaceholders []string
	var args []any
	var duplicateValues []string

	for _, value := range values {
		valueNames = append(valueNames, value.Name)
		valuePlaceholders = append(valuePlaceholders, "?")
		args = append(args, value.Value)
		duplicateValues = append(duplicateValues, fmt.Sprintf("%s = VALUES(%s)", value.Name, value.Name))
	}

	sql = append(sql, fmt.Sprintf("(%s)", strings.Join(valueNames, ", ")))
	sql = append(sql, fmt.Sprintf("VALUES (%s)", strings.Join(valuePlaceholders, ", ")))
	sql = append(sql, "ON DUPLICATE KEY UPDATE")
	sql = append(sql, strings.Join(duplicateValues, ", "))

	return database.DB.Exec(strings.Join(sql, " "), args...)
}

func upsertRowSqlite3(table string, values []structs.QueryValue) (sql.Result, error) {
	sql := []string{fmt.Sprintf(
		"INSERT INTO %s",
		table,
	)}

	var valueNames []string
	var valuePlaceholders []string
	var args []any
	var duplicateValues []string

	for _, value := range values {
		valueNames = append(valueNames, value.Name)
		valuePlaceholders = append(valuePlaceholders, "?")
		args = append(args, value.Value)
		duplicateValues = append(duplicateValues, fmt.Sprintf("%s = excluded.%s", value.Name, value.Name))
	}

	sql = append(sql, fmt.Sprintf("(%s)", strings.Join(valueNames, ", ")))
	sql = append(sql, fmt.Sprintf("VALUES (%s)", strings.Join(valuePlaceholders, ", ")))
	sql = append(sql, "ON CONFLICT DO UPDATE SET")
	sql = append(sql, strings.Join(duplicateValues, ", "))

	return database.DB.Exec(strings.Join(sql, " "), args...)
}
