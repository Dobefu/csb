CREATE TABLE IF NOT EXISTS translations(
  source varchar(128) NOT NULL PRIMARY KEY UNIQUE,
  translation varchar(128) NOT NULL,
  category varchar(255) NOT NULL,
  uid varchar(32) NOT NULL
);
