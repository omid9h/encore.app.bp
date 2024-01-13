// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: insert_user.sql

package repo

import (
	"context"
)

const insertUser = `-- name: InsertUser :one
INSERT INTO users (email, password)
VALUES ($1, $2)
RETURNING id, email, password, roles, active, verified, created_at
`

type InsertUserParams struct {
	Email    string
	Password string
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, insertUser, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Roles,
		&i.Active,
		&i.Verified,
		&i.CreatedAt,
	)
	return i, err
}