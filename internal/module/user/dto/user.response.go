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

func FromDB(user db.User) Response {
	return Response{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}

func FromDBList(users []db.User) []Response {
	var userDTOs []Response
	for _, user := range users {
		userDTOs = append(userDTOs, FromDB(user))
	}
	return userDTOs
}
