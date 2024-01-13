package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var t *jwt.Token

func CreateJWTToken(key, issuer, id, email, roles string) (token string) {

	t = jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss":   issuer,
			"id":    id,
			"email": email,
			"roles": roles,
			"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
			"iat":   time.Now().Unix(),
		})
	token, _ = t.SignedString([]byte(key))
	return token
}
