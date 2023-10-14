package service

import (
	"context"

	"github.com/EdwardKerckhof/gohtmx/config"
	"github.com/EdwardKerckhof/gohtmx/internal/db"
	"github.com/EdwardKerckhof/gohtmx/internal/module/auth/dto"
	"github.com/EdwardKerckhof/gohtmx/pkg/token"
)

type Service interface {
	Register(ctx context.Context, req dto.RegisterRequest) (dto.Response, error)
	Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)
}

type authService struct {
	config     config.Config
	store      db.Store
	tokenMaker token.Maker
}

func New(config config.Config, store db.Store, tokenMaker token.Maker) Service {
	return authService{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
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
	return dto.NewResponse(user), nil
}

func (s authService) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := s.store.FindUserByUsername(ctx, req.Username)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	if err := req.ComparePassword(user.Password); err != nil {
		return dto.LoginResponse{}, err
	}

	accessToken, err := s.tokenMaker.GenerateToken(user.Username, s.config.Auth.AccessTokenExpiration)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.NewLoginResponse(accessToken, user), nil
}
