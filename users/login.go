package users

import (
	"context"
	"errors"

	"github.com/omid9h/encore.app.bp/pkg/hash"
	"github.com/omid9h/encore.app.bp/pkg/token"

	"github.com/invopop/validation"
	"github.com/invopop/validation/is"
)

var secrets struct {
	// JWTSecret is a random key for creating token
	// can be made via `openssl rand -hex 16`
	JWTSecret string
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r LoginInput) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required, validation.Length(5, 100), is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(5, 100)),
	)
}

type LoginOutput struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

// Login checks given email & password and returns a token
//
//encore:api public method=POST path=/user/login
func (s *Service) Login(ctx context.Context, input LoginInput) (output LoginOutput, err error) {
	u, err := s.repo.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return
	}
	if u.Email == "" || !hash.ComparePasswords(u.Password, input.Password) {
		err = ErrUserNotFound
		return
	}
	output.ID = u.ID.String()
	output.Email = u.Email
	output.Token = token.CreateJWTToken(secrets.JWTSecret, s.Issuer, u.ID.String(), u.Email, u.Roles.String)
	return
}

var ErrUserNotFound = errors.New("user not found")
