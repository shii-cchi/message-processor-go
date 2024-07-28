-- +goose Up

CREATE TABLE messages (
    id UUID PRIMARY KEY,
    content TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'unprocessed'
);

-- +goose Down

DROP TABLE messages;