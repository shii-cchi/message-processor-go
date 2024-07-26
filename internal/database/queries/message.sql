-- name: CreateMessage :one
INSERT INTO messages (content)
VALUES ($1)
RETURNING *;

-- name: GetMessagesCount :one
SELECT COUNT(*) FROM messages;

-- name: GetProcessedMessagesCount :one
SELECT COUNT(*) FROM messages
WHERE status = 'processed';