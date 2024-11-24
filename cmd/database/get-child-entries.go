package database

import "github.com/Dobefu/csb/cmd/cs_sdk/structs"

func GetChildEntriesByUid(uid string, locale string) ([]structs.Route, error) {
	rows, err := DB.Query(
		"SELECT * FROM routes WHERE parent = ? AND locale = ?",
		uid,
		locale,
	)

	if err != nil {
		return []structs.Route{}, err
	}

	var results []structs.Route

	for rows.Next() {
		var result structs.Route

		err := rows.Scan(
			&result.Id,
			&result.Uid,
			&result.ContentType,
			&result.Locale,
			&result.Slug,
			&result.Url,
			&result.Parent,
			&result.Published,
		)

		if err != nil {
			continue
		}

		results = append(results, result)
	}

	return results, nil
}
