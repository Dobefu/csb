package query

import (
	"database/sql"

	"github.com/Dobefu/csb/cmd/database"
)

func QueryRaw(sql string) (sql.Result, error) {
	result, err := database.DB.Exec(sql)

	if err != nil {
		return nil, err
	}

	return result, nil
}
