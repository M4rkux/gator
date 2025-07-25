-- +goose Up
CREATE TABLE feeds(
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name TEXT NOT NULL UNIQUE,
  url TEXT NOT NULL,
  user_id UUID NOT NULL,
  FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose down
DROP TABLE seeds;
