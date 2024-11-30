CREATE TABLE IF NOT EXISTS routes(
  id varchar(48) NOT NULL PRIMARY KEY UNIQUE,
  uid varchar(32) NOT NULL,
  contentType varchar(255) NOT NULL,
  locale varchar(16) NOT NULL,
  slug varchar(255) NOT NULL,
  url varchar(255) NOT NULL,
  parent varchar(64),
  published boolean NOT NULL DEFAULT false
);
