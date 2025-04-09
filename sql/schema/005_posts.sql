-- +goose Up
CREATE TABLE posts
(
    id           UUID PRIMARY KEY,
    created_at   TIMESTAMP NOT NULL,
    updated_at   TIMESTAMP NOT NULL DEFAULT NOW(),
    title        TEXT      NOT NULL,
    url          TEXT      NOT NULL,
    description  TEXT      NOT NULL,
    published_at TIMESTAMP NOT NULL,
    feed_id      UUID      NOT NULL,
    CONSTRAINT fk_posts_feeds
        FOREIGN KEY (feed_id)
            REFERENCES feeds (id)
            ON DELETE CASCADE
);

CREATE UNIQUE INDEX idx_uq_posts_url ON posts (url);

-- +goose Down
DROP TABLE posts;