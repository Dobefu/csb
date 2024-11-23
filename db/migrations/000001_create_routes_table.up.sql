CREATE TABLE IF NOT EXISTS routes(
  uid varchar(64) NOT NULL PRIMARY KEY UNIQUE,
  contentType varchar(255) NOT NULL,
  slug varchar(255) NOT NULL,
  url varchar(255) NOT NULL,
  parent varchar(64),
  published boolean NOT NULL DEFAULT 0
);
