package auth

import (
	"github.com/google/uuid"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
)

type Auth struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}

func FromDB(auth db.User) Auth {
	return Auth{
		ID:       auth.ID,
		Username: auth.Username,
		Email:    auth.Email,
	}
}
