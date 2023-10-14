package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token has expired")
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"userId"`
	IssuedAt  time.Time `json:"issuedAt"`
	ExpiredAt time.Time `json:"expiredAt"`
}

func NewPayload(userID uuid.UUID, duration time.Duration) *Payload {
	return &Payload{
		ID:        uuid.New(),
		UserID:    userID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
}

func (c Payload) Valid() error {
	if time.Now().After(c.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
