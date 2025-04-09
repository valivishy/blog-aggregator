-- name: DeleteFeedFollow :exec
delete from feed_follows ff
where ff.user_id = $1
and ff.feed_id in (
    select id from feeds where url = $2
);