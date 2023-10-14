package service

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/EdwardKerckhof/gohtmx/config"
	"github.com/EdwardKerckhof/gohtmx/internal/db"
	"github.com/EdwardKerckhof/gohtmx/internal/module/auth/dto"
	"github.com/EdwardKerckhof/gohtmx/pkg/token"
)

var (
	ErrBlocked        = errors.New("session is blocked")
	ErrIncorrectUser  = errors.New("incorrect user")
	ErrInvalidToken   = errors.New("invalid token")
	ErrExpiredSession = errors.New("session is expired")
)

type Service interface {
	Register(ctx *gin.Context, req dto.RegisterRequest) (dto.Response, error)
	Login(ctx *gin.Context, req dto.LoginRequest) (dto.LoginResponse, error)
	RefreshAccessToken(ctx *gin.Context, req dto.RefreshTokenRequest) (dto.RefreshTokenResponse, error)
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

func (s authService) Register(ctx *gin.Context, req dto.RegisterRequest) (dto.Response, error) {
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

func (s authService) Login(ctx *gin.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := s.store.FindUserByUsername(ctx, req.Username)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	if err := req.ComparePassword(user.Password); err != nil {
		return dto.LoginResponse{}, err
	}

	accessToken, accessPayload, err := s.tokenMaker.GenerateToken(user.ID, s.config.Auth.AccessTokenExpiration)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	refreshToken, refreshPayload, err := s.tokenMaker.GenerateToken(user.ID, s.config.Auth.RefreshTokenExpiration)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	session, err := s.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		UserID:       user.ID,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.NewLoginResponse(user, session.ID, accessToken, accessPayload.ExpiredAt, refreshToken, refreshPayload.ExpiredAt), nil
}

func (s authService) RefreshAccessToken(ctx *gin.Context, req dto.RefreshTokenRequest) (dto.RefreshTokenResponse, error) {
	refreshPayload, err := s.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		return dto.RefreshTokenResponse{}, err
	}

	session, err := s.store.FindSessionById(ctx, refreshPayload.ID)
	if err != nil {
		return dto.RefreshTokenResponse{}, err
	}

	if session.IsBlocked {
		return dto.RefreshTokenResponse{}, ErrBlocked
	}

	if session.UserID != refreshPayload.UserID {
		return dto.RefreshTokenResponse{}, ErrIncorrectUser
	}

	if session.RefreshToken != req.RefreshToken {
		return dto.RefreshTokenResponse{}, ErrInvalidToken
	}

	if time.Now().After(session.ExpiresAt) {
		return dto.RefreshTokenResponse{}, ErrExpiredSession
	}

	accessToken, accessPayload, err := s.tokenMaker.GenerateToken(session.UserID, s.config.Auth.AccessTokenExpiration)
	if err != nil {
		return dto.RefreshTokenResponse{}, err
	}

	return dto.RefreshTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
	}, nil
}
