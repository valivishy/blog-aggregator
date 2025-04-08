-- name: GetFeedFollowsForUser :many
SELECT f.name as feed_name, u.name as user_name
FROM feed_follows ff
INNER JOIN users u on ff.user_id = u.id
INNER JOIN feeds f on ff.feed_id = f.id
WHERE ff.user_id = $1;