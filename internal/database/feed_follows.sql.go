// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: feed_follows.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeedFollow = `-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows(id, user_id, feed_id, created_at, updated_at)
  VALUES ($1, $2, $3, $4, $5)
  RETURNING id, user_id, feed_id, created_at, updated_at
)
SELECT
  iff.id, iff.user_id, iff.feed_id, iff.created_at, iff.updated_at,
  f.name AS feed_name,
  u.name AS user_name
FROM inserted_feed_follow AS iff
INNER JOIN feeds AS f ON iff.feed_id = f.id
INNER JOIN users AS u ON iff.user_id = u.id
`

type CreateFeedFollowParams struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	FeedID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateFeedFollowRow struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	FeedID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	FeedName  string
	UserName  string
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (CreateFeedFollowRow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollow,
		arg.ID,
		arg.UserID,
		arg.FeedID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i CreateFeedFollowRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.FeedID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FeedName,
		&i.UserName,
	)
	return i, err
}

const deleteFeedFollowForUser = `-- name: DeleteFeedFollowForUser :exec
DELETE FROM feed_follows 
WHERE feed_follows.user_id = $1 
AND feed_follows.feed_id = $2
`

type DeleteFeedFollowForUserParams struct {
	UserID uuid.UUID
	FeedID uuid.UUID
}

func (q *Queries) DeleteFeedFollowForUser(ctx context.Context, arg DeleteFeedFollowForUserParams) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollowForUser, arg.UserID, arg.FeedID)
	return err
}

const getFeedFollowsForUser = `-- name: GetFeedFollowsForUser :many
SELECT 
  ff.id, ff.user_id, ff.feed_id, ff.created_at, ff.updated_at,
  f.name AS feed_name,
  u.name AS user_name
FROM feed_follows AS ff
INNER JOIN feeds AS f ON ff.feed_id = f.id
INNER JOIN users AS u ON ff.user_id = u.id
WHERE ff.user_id = $1
`

type GetFeedFollowsForUserRow struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	FeedID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	FeedName  string
	UserName  string
}

func (q *Queries) GetFeedFollowsForUser(ctx context.Context, userID uuid.UUID) ([]GetFeedFollowsForUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollowsForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFeedFollowsForUserRow
	for rows.Next() {
		var i GetFeedFollowsForUserRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.FeedID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.FeedName,
			&i.UserName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
