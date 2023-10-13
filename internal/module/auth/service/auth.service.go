package service

import (
	"context"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
	"github.com/EdwardKerckhof/gohtmx/internal/module/auth/dto"
)

type Service interface {
	Register(ctx context.Context, req dto.RegisterRequest) (dto.Response, error)
	Login(ctx context.Context, req dto.LoginRequest) (dto.Response, error)
}

type authService struct {
	store db.Store
}

func New(store db.Store) Service {
	return authService{
		store: store,
	}
}

func (s authService) Register(ctx context.Context, req dto.RegisterRequest) (dto.Response, error) {
	hashedPassword, err := req.HashPassword()
	if err != nil {
		return dto.Response{}, err
	}

	arg := db.CreateUserParams{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
	}
	user, err := s.store.CreateUser(ctx, arg)
	if err != nil {
		return dto.Response{}, err
	}
	return dto.FromDB(user), nil
}

func (s authService) Login(ctx context.Context, req dto.LoginRequest) (dto.Response, error) {
	panic("unimplemented")
}
