// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: insert_token.sql

package repo

import (
	"context"

	"github.com/google/uuid"
)

const insertToken = `-- name: InsertToken :exec
INSERT INTO tokens (user_id, token)
VALUES ($1, $2) ON CONFLICT (user_id) DO
UPDATE
SET token = EXCLUDED.token
`

type InsertTokenParams struct {
	UserID uuid.UUID
	Token  string
}

func (q *Queries) InsertToken(ctx context.Context, arg InsertTokenParams) error {
	_, err := q.db.ExecContext(ctx, insertToken, arg.UserID, arg.Token)
	return err
}