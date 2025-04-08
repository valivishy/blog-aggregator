-- name: ListFeeds :many
SELECT f.name as feed_name, f.url as feed_url, u.name as user_name
FROM feeds f
INNER JOIN users u on f.user_id = u.id;