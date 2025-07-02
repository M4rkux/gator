-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows(id, user_id, feed_id, created_at, updated_at)
  VALUES ($1, $2, $3, $4, $5)
  RETURNING *
)
SELECT
  iff.*,
  f.name AS feed_name,
  u.name AS user_name
FROM inserted_feed_follow AS iff
INNER JOIN feeds AS f ON iff.feed_id = f.id
INNER JOIN users AS u ON iff.user_id = u.id;

-- name: GetFeedFollowsForUser :many
SELECT 
  ff.*,
  f.name AS feed_name,
  u.name AS user_name
FROM feed_follows AS ff
INNER JOIN feeds AS f ON ff.feed_id = f.id
INNER JOIN users AS u ON ff.user_id = u.id
WHERE ff.user_id = $1;

