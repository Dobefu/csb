package database

func SetRoute(
	id string,
	uid string,
	contentType string,
	locale string,
	slug string,
	url string,
	parent string,
	published bool,
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
		id,
		uid,
		contentType,
		locale,
		slug,
		url,
		parent,
		published,
	)

	if err != nil {
		return err
	}
	return nil
}
