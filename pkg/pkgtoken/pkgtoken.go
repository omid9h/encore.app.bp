package pkgtoken

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var t *jwt.Token

type Claims struct {
	Issuer    string    `json:"issuer"`
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Roles     string    `json:"roles"`
	ExpiresAt time.Time `json:"expires_at"`
	IssuedAt  time.Time `json:"issued_at"`
}

func MapClaimsToStruct(claims jwt.MapClaims) (*Claims, error) {
	var customClaims Claims

	if issuer, ok := claims["iss"].(string); ok {
		customClaims.Issuer = issuer
	} else {
		return nil, fmt.Errorf("issuer claim not found or not a string")
	}

	if id, ok := claims["id"].(string); ok {
		customClaims.ID = id
	} else {
		return nil, fmt.Errorf("iD claim not found or not a string")
	}

	if email, ok := claims["email"].(string); ok {
		customClaims.Email = email
	} else {
		return nil, fmt.Errorf("email claim not found or not a string")
	}

	if roles, ok := claims["roles"].(string); ok {
		customClaims.Roles = roles
	} else {
		return nil, fmt.Errorf("roles claim not found or not a string")
	}

	if exp, ok := claims["exp"].(float64); ok {
		customClaims.ExpiresAt = time.Unix(int64(exp), 0)
	} else {
		return nil, fmt.Errorf("expiresAt claim not found or not a float64")
	}

	if iat, ok := claims["iat"].(float64); ok {
		customClaims.IssuedAt = time.Unix(int64(iat), 0)
	} else {
		return nil, fmt.Errorf("issuedAt claim not found or not a float64")
	}

	return &customClaims, nil
}

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

func DecodeJWTToken(tokenString string, key string) (claims jwt.MapClaims, err error) {
	t, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, ErrInvalidToken
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed to extract claims")
	}
	return claims, nil
}

var ErrInvalidToken = errors.New("token is not valid")
var ErrExpiredToken = errors.New("token has expired")
