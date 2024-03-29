// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package repo

import (
	"database/sql"

	"github.com/google/uuid"
)

type Token struct {
	UserID uuid.UUID
	Token  string
}

type User struct {
	ID        uuid.UUID
	Email     string
	Password  string
	Roles     sql.NullString
	Active    sql.NullBool
	Verified  sql.NullBool
	CreatedAt sql.NullTime
}
