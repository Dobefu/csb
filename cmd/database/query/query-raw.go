package query

import (
	"database/sql"

	"github.com/Dobefu/csb/cmd/database"
)

func QueryRaw(sql string) (sql.Result, error) {
	return database.DB.Exec(sql)
}
