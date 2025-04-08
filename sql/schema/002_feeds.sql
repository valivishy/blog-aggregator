-- +goose Up
CREATE TABLE feeds
(
    id         UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    name       TEXT      NOT NULL,
    url        TEXT      NOT NULL,
    user_id    UUID      NOT NULL,
    CONSTRAINT fk_feeds_users
        FOREIGN KEY (user_id)
            REFERENCES users (id)
            ON DELETE CASCADE
);

CREATE UNIQUE INDEX unq_url ON feeds (url);

-- +goose Down
DROP TABLE feeds;