package jwt

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/require"

	"github.com/EdwardKerckhof/gohtmx/pkg/token"
)

func TestJWTMaker(t *testing.T) {
	secret := "12345678901234567890123456789012"
	maker, err := NewMaker(secret)
	require.NoError(t, err)

	username := "test"
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	jwt, err := maker.GenerateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, jwt)

	payload, err := maker.VerifyToken(jwt)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredJWT(t *testing.T) {
	secret := "12345678901234567890123456789012"
	maker, err := NewMaker(secret)
	require.NoError(t, err)

	username := "test"
	jwt, err := maker.GenerateToken(username, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, jwt)

	payload, err := maker.VerifyToken(jwt)
	require.Error(t, err)
	require.EqualError(t, err, token.ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TextInvalidJWTAlgNone(t *testing.T) {
	payload := token.NewPayload("test", time.Minute)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	tok, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	secret := "12345678901234567890123456789012"
	maker, err := NewMaker(secret)
	require.NoError(t, err)

	payload, err = maker.VerifyToken(tok)
	require.Error(t, err)
	require.EqualError(t, err, token.ErrInvalidToken.Error())
	require.Nil(t, payload)
}
