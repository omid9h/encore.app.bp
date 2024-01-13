package users

import (
	"context"
	"time"

	"github.com/omid9h/encore.app.bp/pkg/pkghash"

	"github.com/google/uuid"
	"github.com/invopop/validation"
	"github.com/invopop/validation/is"
	"github.com/omid9h/encore.app.bp/users/repo"
)

type RegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r RegisterInput) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required, validation.Length(5, 100), is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(5, 100)),
	)
}

type RegisterOutput struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Roles     string    `json:"roles"`
	Active    bool      `json:"active"`
	Verified  bool      `json:"verified"`
	CreatedAt time.Time `json:"created_at"`
}

// Register inserts a new user with given email & password
// and other default values.
//
//encore:api public method=POST path=/user/register
func (s *Service) Register(ctx context.Context, input RegisterInput) (output RegisterOutput, err error) {
	input.Password = pkghash.HashPassword(input.Password)
	u, err := s.repo.InsertUser(ctx, repo.InsertUserParams{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return
	}
	output.ID = u.ID
	output.Email = u.Email
	output.Roles = u.Roles.String
	output.Active = u.Active.Bool
	output.Verified = u.Verified.Bool
	output.CreatedAt = u.CreatedAt.Time
	return
}
