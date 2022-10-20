CREATE TABLE food (
  id integer PRIMARY KEY AUTOINCREMENT,
  name text NOT NULL,
  unit text NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
);
