package query

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/database/structs"
	"github.com/Dobefu/csb/cmd/logger"
)

func Upsert(table string, values []structs.QueryValue) error {
	var (
		sql  string
		args []any
	)

	dbType := os.Getenv("DB_TYPE")

	switch dbType {
	case "mysql":
		sql, args = upsertRowMysql(table, values)
	case "sqlite3":
		sql, args = upsertRowSqlite3(table, values)
	case "postgres":
		sql, args = upsertRowPostgres(table, values)
	default:
		logger.Fatal(
			"The database type %s has no corresponding QueryRow function",
			dbType,
		)
	}

	_, err := database.DB.Exec(sql, args...)

	return err
}

func upsertRowMysql(table string, values []structs.QueryValue) (string, []any) {
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

	return strings.Join(sql, " "), args
}

func upsertRowSqlite3(table string, values []structs.QueryValue) (string, []any) {
	sql, args := upsertRowMysql(table, values)
	sql = strings.Replace(sql, "DUPLICATE KEY UPDATE", "CONFLICT DO UPDATE SET", 1)

	sql = regexp.MustCompile(`VALUES\((.+?)\)`).ReplaceAllString(sql, "excluded.$1")

	return sql, args
}

func upsertRowPostgres(table string, values []structs.QueryValue) (string, []any) {
	sql, args := upsertRowMysql(table, values)
	sql = strings.Replace(
		sql,
		"DUPLICATE KEY UPDATE",
		fmt.Sprintf("CONFLICT ON CONSTRAINT %s_pkey DO UPDATE SET", table),
		1,
	)

	sql = regexp.MustCompile(`VALUES\((.+?)\)`).ReplaceAllString(sql, "EXCLUDED.$1")

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
