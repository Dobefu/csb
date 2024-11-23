CREATE TABLE IF NOT EXISTS state(
  id int NOT NULL PRIMARY KEY UNIQUE,
  name varchar(255) NOT NULL,
  value varchar(255) NOT NULL
);
