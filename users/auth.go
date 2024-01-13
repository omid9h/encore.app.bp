package users

import (
	"context"

	"encore.dev/beta/auth"
)

type UserData struct {
	Email    string `json:"email"`
	Roles    string `json:"roles"`
	Active   bool   `json:"active"`
	Verified bool   `json:"verified"`
}

//encore:authhandler
func (s *Service) AuthHandler(ctx context.Context, token string) (auth.UID, *UserData, error) {
	// Validate the token and look up the user id and user data,
	return auth.UID(""), nil, nil
}
