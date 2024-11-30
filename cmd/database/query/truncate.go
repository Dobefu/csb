package query

import (
	"fmt"
	"os"

	"github.com/Dobefu/csb/cmd/database"
)

func Truncate(table string) error {
	switch os.Getenv("DB_TYPE") {
	case "mysql":
		return truncateMysql(table)
	case "sqlite3":
		return truncateSqlite3(table)
	default:
		return nil
	}
}

func truncateMysql(table string) error {
	sql := fmt.Sprintf("TRUNCATE %s", table)
	_, err := database.DB.Exec(sql)

	if err != nil {
		return err
	}

	return nil
}

func truncateSqlite3(table string) error {
	sql := fmt.Sprintf("DELETE FROM %s", table)
	_, err := database.DB.Exec(sql)

	if err != nil {
		return err
	}

	return nil
}
