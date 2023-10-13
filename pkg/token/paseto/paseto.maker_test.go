package paseto

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/EdwardKerckhof/gohtmx/pkg/token"
)

func TestPasetoMaker(t *testing.T) {
	secret := "12345678901234567890123456789012"
	maker, err := NewPasetoMaker(secret)
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

func TestExpiredPaseto(t *testing.T) {
	secret := "12345678901234567890123456789012"
	maker, err := NewPasetoMaker(secret)
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
