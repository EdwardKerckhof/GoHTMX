package dto

import (
	"github.com/google/uuid"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
)

type Response struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}

type LoginResponse struct {
	AccessToken string   `json:"accessToken"`
	User        Response `json:"user"`
}

func NewResponse(dbAuth db.User) Response {
	return Response{
		ID:       dbAuth.ID,
		Username: dbAuth.Username,
		Email:    dbAuth.Email,
	}
}

func NewLoginResponse(accessToken string, dbAuth db.User) LoginResponse {
	return LoginResponse{
		AccessToken: accessToken,
		User:        NewResponse(dbAuth),
	}
}
