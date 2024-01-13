package users

import (
	"context"
	"time"

	"encore.dev/beta/auth"
	"github.com/google/uuid"
	"github.com/omid9h/encore.app.bp/pkg/pkgtoken"
)

type UserData struct {
	UID       auth.UID  `json:"uid"`
	Email     string    `json:"email"`
	Roles     string    `json:"roles"`
	ExpiresAt time.Time `json:"exp"`
}

//encore:authhandler
func (s *Service) AuthHandler(ctx context.Context, token string) (uid auth.UID, udata *UserData, err error) {
	udata = &UserData{}
	// decode token
	claimsMap, err := pkgtoken.DecodeJWTToken(token, secrets.JWTSecret)
	if err != nil {
		return
	}
	claimsStruct, err := pkgtoken.MapClaimsToStruct(claimsMap)
	if err != nil {
		return
	}
	uid = auth.UID(claimsStruct.ID)
	udata.UID = uid
	udata.Email = claimsStruct.Email
	udata.Roles = claimsStruct.Roles
	udata.ExpiresAt = claimsStruct.ExpiresAt
	// verify against repo
	t, err := s.repo.GetTokenByUserID(ctx, uuid.Must(uuid.Parse(string(uid))))
	if err != nil {
		err = pkgtoken.ErrInvalidToken
		return
	}
	if t.Token != token {
		err = pkgtoken.ErrInvalidToken
		return
	}
	// check if token has expired
	if udata.ExpiresAt.Before(time.Now()) {
		err = pkgtoken.ErrExpiredToken
		return
	}
	return
}
