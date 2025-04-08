-- name: CreateFeedFollow :one
WITH feed_follow as (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
           $1,
           $2,
           $3,
           $4,
           $5
       )
    RETURNING *
)
SELECT feed_follow.*, u.name as linked_user, f.name as linked_feed
FROM feed_follow
INNER JOIN users u on feed_follow.user_id = u.id
INNER JOIN feeds f on feed_follow.feed_id = f.id;


