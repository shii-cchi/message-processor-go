-- name: CreateMessage :one
INSERT INTO messages (id, content)
VALUES ($1, $2)
RETURNING *;

-- name: GetMessagesCount :one
SELECT COUNT(*) FROM messages;

-- name: UpdateMessageStatus :exec
UPDATE messages
SET status = $2
WHERE id = $1;

-- name: GetProcessedMessagesCount :one
SELECT COUNT(*) FROM messages
WHERE status = 'processed';