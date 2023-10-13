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

func FromDB(dbAuth db.User) Response {
	return Response{
		ID:       dbAuth.ID,
		Username: dbAuth.Username,
		Email:    dbAuth.Email,
	}
}
