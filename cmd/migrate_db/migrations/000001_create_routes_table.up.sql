CREATE TABLE IF NOT EXISTS routes(
  id varchar(48) NOT NULL PRIMARY KEY UNIQUE,
  uid varchar(32) NOT NULL,
  title varchar(255) NOT NULL,
  content_type varchar(255) NOT NULL,
  locale varchar(16) NOT NULL,
  slug varchar(255) NOT NULL,
  url varchar(255) NOT NULL,
  parent varchar(64),
  version int DEFAULT 0,
  updated_at timestamp NOT NULL,
  exclude_sitemap boolean NOT NULL DEFAULT false,
  published boolean NOT NULL DEFAULT false
);
