package auth

import (
	"context"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
	authRequest "github.com/EdwardKerckhof/gohtmx/internal/dto/request/auth"
	authResponse "github.com/EdwardKerckhof/gohtmx/internal/dto/response/auth"
)

type Service interface {
	Register(ctx context.Context, req authRequest.RegisterRequest) (authResponse.Auth, error)
	Login(ctx context.Context, req authRequest.LoginRequest) (authResponse.Auth, error)
}

type authService struct {
	store db.Store
}

func New(store db.Store) Service {
	return authService{
		store: store,
	}
}

func (s authService) Register(ctx context.Context, req authRequest.RegisterRequest) (authResponse.Auth, error) {
	hashedPassword, err := req.HashPassword()
	if err != nil {
		return authResponse.Auth{}, err
	}

	arg := db.CreateUserParams{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
	}
	user, err := s.store.CreateUser(ctx, arg)
	if err != nil {
		return authResponse.Auth{}, err
	}
	return authResponse.FromDB(user), nil
}

func (s authService) Login(ctx context.Context, req authRequest.LoginRequest) (authResponse.Auth, error) {
	panic("unimplemented")
}
