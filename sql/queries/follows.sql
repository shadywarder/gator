-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
  VALUES ($1, $2, $3, $4, $5) 
  RETURNING *
)
SELECT inserted_feed_follow.*, f.name AS feed_name, u.name AS user_name
FROM inserted_feed_follow
JOIN users u ON user_id = u.id
JOIN feeds f ON feed_id = f.id;

-- name: GetFeedFollowsForUser :many
SELECT u.name AS user_name, f.name AS feed_name
FROM feed_follows ff
JOIN users u ON ff.user_id = u.id
JOIN feeds f ON ff.feed_id = f.id
WHERE u.name = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows ff
USING users u, feeds f
WHERE ff.user_id = u.id AND ff.feed_id = f.id AND u.name = $1 AND f.url = $2;