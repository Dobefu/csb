CREATE TABLE IF NOT EXISTS assets(
  id varchar(48) NOT NULL PRIMARY KEY UNIQUE,
  uid varchar(32) NOT NULL,
  title varchar(255) NOT NULL,
  content_type varchar(255) NOT NULL,
  locale varchar(16) NOT NULL,
  url varchar(255) NOT NULL,
  parent varchar(64),
  height int NOT NULL,
  width int NOT NULL,
  updated_at timestamp NOT NULL,
  published boolean NOT NULL DEFAULT false
);
