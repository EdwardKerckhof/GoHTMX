package dto

import (
	"time"

	"github.com/google/uuid"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
)

type Response struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}

type LoginResponse struct {
	SessionID             uuid.UUID `json:"sessionId"`
	AccessToken           string    `json:"accessToken"`
	AccessTokenExpiresAt  time.Time `json:"accessTokenExpiresAt"`
	RefreshToken          string    `json:"refreshToken"`
	RefreshTokenExpiresAt time.Time `json:"refreshTokenExpiresAt"`
	User                  Response  `json:"user"`
}

type RefreshTokenResponse struct {
	AccessToken          string    `json:"accessToken"`
	AccessTokenExpiresAt time.Time `json:"accessTokenExpiresAt"`
}

func NewResponse(dbAuth db.User) Response {
	return Response{
		ID:       dbAuth.ID,
		Username: dbAuth.Username,
		Email:    dbAuth.Email,
	}
}

func NewLoginResponse(dbAuth db.User, sessionID uuid.UUID, accessToken string, accessTokenExpiresAt time.Time, refreshToken string, refreshTokenExpiresAt time.Time) LoginResponse {
	return LoginResponse{
		SessionID:             sessionID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessTokenExpiresAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshTokenExpiresAt,
		User:                  NewResponse(dbAuth),
	}
}
