CREATE TABLE IF NOT EXISTS translations(
  id varchar(48) NOT NULL PRIMARY KEY UNIQUE,
  uid varchar(32) NOT NULL,
  source varchar(128) NOT NULL,
  translation varchar(128) NOT NULL,
  category varchar(255) NOT NULL,
  locale varchar(16) NOT NULL
);
