-- +goose Up
CREATE TABLE feed_follows(
  id UUID PRIMARY KEY,
  user_id UUID NOT NULL,
  feed_id UUID NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY(feed_id) REFERENCES feeds(id) ON DELETE CASCADE,
  CONSTRAINT unique_user_feed UNIQUE (user_id, feed_id)
);

-- +goose down
DROP TABLE feed_follows;
