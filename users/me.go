package users

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type MeOutput struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Roles     string    `json:"roles"`
	Active    bool      `json:"active"`
	Verified  bool      `json:"verified"`
	CreatedAt time.Time `json:"created_at"`
}

// Me returns logged-in user's info
//
//encore:api auth method=GET path=/user/me
func (s *Service) Me(ctx context.Context) (output MeOutput, err error) {
	return
}
