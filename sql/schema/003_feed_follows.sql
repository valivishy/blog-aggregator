-- +goose Up
CREATE TABLE feed_follows
(
    id         UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    user_id    UUID      NOT NULL,
    feed_id    UUID      NOT NULL,
    CONSTRAINT fk_feed_follows_users
        FOREIGN KEY (user_id)
            REFERENCES users (id)
            ON DELETE CASCADE,
    CONSTRAINT fk_feed_follows_feeds
        FOREIGN KEY (feed_id)
            REFERENCES feeds (id)
            ON DELETE CASCADE
);

CREATE UNIQUE INDEX unq_feed_follows_feed_user ON feed_follows (user_id, feed_id);

-- +goose Down
DROP TABLE feed_follows;