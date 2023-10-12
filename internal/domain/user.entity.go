package domain

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseEntity
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

// TODO: Add these functions to service add unit test (backend #17)
func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
