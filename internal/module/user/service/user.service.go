package service

import (
	"context"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
	"github.com/EdwardKerckhof/gohtmx/internal/module/user/dto"
	"github.com/EdwardKerckhof/gohtmx/pkg/request"
)

type Service interface {
	Count(ctx context.Context) (int64, error)
	FindAll(ctx context.Context, req dto.FindAllRequest) ([]dto.Response, error)
	FindAllWithCount(ctx context.Context, req dto.FindAllRequest) ([]dto.Response, int64, error)
	FindById(ctx context.Context, id request.IDRequest) (dto.Response, error)
}

type userService struct {
	store db.Store
}

func New(store db.Store) Service {
	return userService{
		store: store,
	}
}

func (s userService) Count(ctx context.Context) (int64, error) {
	return s.store.CountUsers(ctx)
}

func (s userService) FindAll(ctx context.Context, req dto.FindAllRequest) ([]dto.Response, error) {
	arg := db.FindAllUsersParams{
		Limit:  req.Size,
		Offset: (req.Page - 1) * req.Size,
	}

	users, err := s.store.FindAllUsers(ctx, arg)
	if err != nil {
		return []dto.Response{}, err
	}
	return dto.NewResponseList(users), nil
}

func (s userService) FindAllWithCount(ctx context.Context, req dto.FindAllRequest) ([]dto.Response, int64, error) {
	users, err := s.FindAll(ctx, req)
	if err != nil {
		return []dto.Response{}, 0, err
	}

	count, err := s.Count(ctx)
	if err != nil {
		return []dto.Response{}, 0, err
	}
	return users, count, nil
}

func (s userService) FindById(ctx context.Context, req request.IDRequest) (dto.Response, error) {
	id, err := req.ParseID()
	if err != nil {
		return dto.Response{}, err
	}

	user, err := s.store.FindUserById(ctx, id)
	if err != nil {
		return dto.Response{}, err
	}
	resp := dto.NewResponse(user)
	return resp, nil
}
