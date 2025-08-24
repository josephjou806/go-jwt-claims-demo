package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenManager interface {
	Generate(subject string, ttl time.Duration) (string, error)
	Validate(tokenStr string) (*jwt.RegisteredClaims, error)
}

type jwtTokenManager struct {
	secret []byte
}

func NewTokenManager(secret string) TokenManager {
	return &jwtTokenManager{secret: []byte(secret)}
}

func (m *jwtTokenManager) Generate(subject string, ttl time.Duration) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   subject,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(m.secret)
}

func (m *jwtTokenManager) Validate(tokenStr string) (*jwt.RegisteredClaims, error) {
	t, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected signing method")
		}
		return m.secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := t.Claims.(*jwt.RegisteredClaims); ok && t.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
