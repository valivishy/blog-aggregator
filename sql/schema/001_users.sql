-- +goose Up
CREATE TABLE users
(
    id         UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    name       TEXT      NOT NULL
);

CREATE UNIQUE INDEX unq_name ON users (name);

-- +goose Down
DROP TABLE users;