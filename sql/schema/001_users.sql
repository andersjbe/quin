-- +goose Up
CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  first_name varchar(127) NOT NULL,
  last_name varchar(127) NOT NULL,
  email varchar(127) NOT NULL UNIQUE,
  password varchar(127) NOT NULL
);

-- +goose Down
DROP TABLE users;
