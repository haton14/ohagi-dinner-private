CREATE TABLE food (
  id integer PRIMARY KEY AUTOINCREMENT,
  name text NOT NULL,
  unit text NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE dinner (
  id integer PRIMARY KEY AUTOINCREMENT,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
)

CREATE TABLE dinner_menu (
  dinner_id integer NOT NULL,
  food_id integer NOT NULL,
  quantity real NOT NULL
)
