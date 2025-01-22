package api

import (
	"fmt"

	"github.com/Dobefu/csb/cmd/api/structs"
	"github.com/Dobefu/csb/cmd/database/query"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
)

var queryQueryRows = query.QueryRows

func GetTranslations(locale string) (map[string]interface{}, error) {
	translations := map[string]interface{}{}

	rows, err := queryQueryRows(
		"translations",
		[]string{"source", "translation", "category"},
		[]db_structs.QueryWhere{
			{
				Name:  "locale",
				Value: locale,
			},
		},
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var result structs.Translation

		err := rows.Scan(
			&result.Source,
			&result.Translation,
			&result.Category,
		)

		if err != nil {
			continue
		}

		id := fmt.Sprintf("%s.%s", result.Category, result.Source)
		translations[id] = result.Translation
	}

	return translations, nil
}
