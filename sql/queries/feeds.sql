-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
values(
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT f.*, u.name AS username FROM feeds f 
INNER JOIN users AS u ON f.user_id = u.id;

-- name: GetFeedByURL :one
SELECT f.*, u.name AS username FROM feeds f
INNER JOIN users AS u ON f.user_id = u.id
WHERE f.url = $1;
