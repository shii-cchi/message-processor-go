// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: message.sql

package database

import (
	"context"
)

const createMessage = `-- name: CreateMessage :one
INSERT INTO messages (content)
VALUES ($1)
RETURNING id, content, status
`

func (q *Queries) CreateMessage(ctx context.Context, content string) (Message, error) {
	row := q.db.QueryRowContext(ctx, createMessage, content)
	var i Message
	err := row.Scan(&i.ID, &i.Content, &i.Status)
	return i, err
}

const getMessagesCount = `-- name: GetMessagesCount :one
SELECT COUNT(*) FROM messages
`

func (q *Queries) GetMessagesCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getMessagesCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getProcessedMessagesCount = `-- name: GetProcessedMessagesCount :one
SELECT COUNT(*) FROM messages
WHERE status = 'processed'
`

func (q *Queries) GetProcessedMessagesCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getProcessedMessagesCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}
