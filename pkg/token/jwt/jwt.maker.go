package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/EdwardKerckhof/gohtmx/pkg/token"
	"github.com/golang-jwt/jwt"
)

const (
	minSecretKeySize = 32
)

type JWTMaker struct {
	secretKey string
}

func NewMaker(secretKey string) (token.Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

func (m *JWTMaker) GenerateToken(username string, duration time.Duration) (string, error) {
	payload := token.NewPayload(username, duration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(m.secretKey))
}

func (m *JWTMaker) VerifyToken(t string) (*token.Payload, error) {
	keyFunc := func(jwtToken *jwt.Token) (interface{}, error) {
		_, ok := jwtToken.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, token.ErrInvalidToken
		}
		return []byte(m.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(t, &token.Payload{}, keyFunc)
	if err != nil {
		validationErr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(validationErr.Inner, token.ErrExpiredToken) {
			return nil, token.ErrExpiredToken
		}
		return nil, token.ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*token.Payload)
	if !ok {
		return nil, token.ErrInvalidToken
	}

	return payload, nil
}
