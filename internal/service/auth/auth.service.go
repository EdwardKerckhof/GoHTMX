package auth

import (
	"context"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
	authRequest "github.com/EdwardKerckhof/gohtmx/internal/dto/request/auth"
	userModel "github.com/EdwardKerckhof/gohtmx/internal/model/user"
)

type Service interface {
	Register(ctx context.Context, req authRequest.RegisterRequest) (userModel.User, error)
	Login(ctx context.Context, req authRequest.LoginRequest) (userModel.User, error)
}

type authService struct {
	store db.Store
}

func New(store db.Store) Service {
	return authService{
		store: store,
	}
}

func (s authService) Register(ctx context.Context, req authRequest.RegisterRequest) (userModel.User, error) {
	hashedPassword, err := req.HashPassword()
	if err != nil {
		return userModel.User{}, err
	}

	arg := db.CreateUserParams{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
	}
	user, err := s.store.CreateUser(ctx, arg)
	if err != nil {
		return userModel.User{}, err
	}
	return userModel.FromDB(user), nil
}

func (s authService) Login(ctx context.Context, req authRequest.LoginRequest) (userModel.User, error) {
	panic("unimplemented")
}
