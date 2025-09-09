package query

import (
	"os"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/logger"
)

func Truncate(table string) error {
	dbType := os.Getenv("DB_TYPE")

	switch dbType {
	case "mysql":
		return truncateMysql(table)
	case "sqlite3":
		return truncateSqlite3(table)
	case "postgres":
		return truncatePostgres(table)
	default:
		logger.Fatal(
			"The database type %s has no corresponding Truncate function",
			dbType,
		)

		return nil
	}
}

func truncateMysql(table string) error {
	_, err := database.DB.Exec("TRUNCATE ?", table)

	if err != nil {
		return err
	}

	return nil
}

func truncateSqlite3(table string) error {
	_, err := database.DB.Exec("DELETE FROM ?", table)

	if err != nil {
		return err
	}

	return nil
}

func truncatePostgres(table string) error {
	return truncateMysql(table)
}
