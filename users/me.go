package users

import (
	"context"
	"time"

	"encore.dev/beta/auth"
)

type MeOutput struct {
	UID       auth.UID  `json:"uid"`
	Email     string    `json:"email"`
	Roles     string    `json:"roles"`
	ExpiresAt time.Time `json:"exp"`
}

// Me returns logged-in user's info
//
//encore:api auth method=GET path=/user/me
func (s *Service) Me(ctx context.Context) (output MeOutput, err error) {
	udata := auth.Data().(*UserData)
	output.UID = udata.UID
	output.Email = udata.Email
	output.Roles = udata.Roles
	output.ExpiresAt = udata.ExpiresAt
	return
}
