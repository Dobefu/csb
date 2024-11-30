package query

import (
	"fmt"
	"os"
	"strings"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/database/structs"
	"github.com/Dobefu/csb/cmd/logger"
)

func Insert(table string, values []structs.QueryValue) error {
	var (
		sql  string
		args []any
	)

	dbType := os.Getenv("DB_TYPE")

	switch dbType {
	case "mysql":
		sql, args = insertRowMysql(table, values)
	case "sqlite3":
		sql, args = insertRowSqlite3(table, values)
	case "postgres":
		sql, args = insertRowPostgres(table, values)
	default:
		logger.Fatal(
			"The database type %s has no corresponding Insert function",
			dbType,
		)

		return nil
	}

	_, err := database.DB.Exec(sql, args...)

	return err
}

func insertRowMysql(table string, values []structs.QueryValue) (string, []any) {
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

	return strings.Join(sql, " "), args
}

func insertRowSqlite3(table string, values []structs.QueryValue) (string, []any) {
	sql, args := insertRowMysql(table, values)

	return sql, args
}

func insertRowPostgres(table string, values []structs.QueryValue) (string, []any) {
	sql, args := insertRowMysql(table, values)

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
