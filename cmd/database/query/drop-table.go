package query

import (
	"fmt"

	"github.com/Dobefu/csb/cmd/database"
)

func DropTable(table string) error {
	sql := fmt.Sprintf("DROP TABLE %s", table)
	_, err := database.DB.Exec(sql)

	if err != nil {
		return err
	}

	return nil
}
