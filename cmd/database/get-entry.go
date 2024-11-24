package database

import "github.com/Dobefu/csb/cmd/cs_sdk/structs"

func GetEntryByParentUid(uid string, locale string) (structs.Route, error) {
	row := DB.QueryRow(
		"SELECT * FROM routes WHERE parent = ? AND locale = ?",
		uid,
		locale,
	)

	var result structs.Route

	err := row.Scan(
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
		return structs.Route{}, err
	}

	return result, nil
}
