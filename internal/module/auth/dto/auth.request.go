package dto

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string `json:"username" form:"username" binding:"required,min=1,max=255,alphanum"`
	Password string `json:"password" form:"password" binding:"required,min=5,max=255"`
	Email    string `json:"email" form:"email" binding:"required,email"`
}

type LoginRequest struct {
	Username string `json:"username" form:"username" binding:"required,min=1,max=255,alphanum"`
	Password string `json:"password" form:"password" binding:"required,min=5,max=255"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" form:"refreshToken" binding:"required"`
}

func (r *RegisterRequest) HashPassword() (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hash), nil
}

func (r *LoginRequest) ComparePassword(hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(r.Password))
}
