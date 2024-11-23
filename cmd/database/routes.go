package database

import (
	"fmt"

	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
)

func SetRoute(
	route structs.Route,
) error {
	_, err := DB.Exec(
		`INSERT INTO routes
      (id, uid, contentType, locale, slug, url, parent, published)
      VALUES (?, ?, ?, ?, ?, ?, ?, ?)
      ON DUPLICATE KEY UPDATE
      id = VALUES(id),
      uid = VALUES(uid),
      contentType = VALUES(contentType),
      locale = VALUES(locale),
      slug = VALUES(slug),
      url = VALUES(url),
      parent = VALUES(parent),
      published = VALUES(published)
  `,
		fmt.Sprintf("%s%s", route.Uid, route.Locale),
		route.Uid,
		route.ContentType,
		route.Locale,
		route.Slug,
		route.Url,
		route.Parent,
		route.Published,
	)

	if err != nil {
		return err
	}

	return nil
}
