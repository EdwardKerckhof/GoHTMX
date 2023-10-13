package user

import (
	"github.com/google/uuid"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}

func FromDBUser(user db.User) User {
	return User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}

func FromDBUsers(users []db.User) []User {
	var userDTOs []User
	for _, user := range users {
		userDTOs = append(userDTOs, FromDBUser(user))
	}
	return userDTOs
}
