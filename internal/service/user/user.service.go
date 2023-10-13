package user

import (
	"context"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
	"github.com/EdwardKerckhof/gohtmx/internal/dto/request"
	userRequest "github.com/EdwardKerckhof/gohtmx/internal/dto/request/user"
	userResponse "github.com/EdwardKerckhof/gohtmx/internal/dto/response/user"
)

type Service interface {
	Count(ctx context.Context) (int64, error)
	FindAll(ctx context.Context, req userRequest.FindAllRequest) ([]userResponse.User, error)
	FindById(ctx context.Context, id request.IDRequest) (userResponse.User, error)
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

func (s userService) FindAll(ctx context.Context, req userRequest.FindAllRequest) ([]userResponse.User, error) {
	arg := db.FindAllUsersParams{
		Limit:  req.Size,
		Offset: (req.Page - 1) * req.Size,
	}

	users, err := s.store.FindAllUsers(ctx, arg)
	if err != nil {
		return []userResponse.User{}, err
	}
	resp := userResponse.FromDBList(users)
	return resp, nil
}

func (s userService) FindById(ctx context.Context, req request.IDRequest) (userResponse.User, error) {
	id, err := req.ParseID()
	if err != nil {
		return userResponse.User{}, err
	}

	user, err := s.store.FindUserById(ctx, id)
	if err != nil {
		return userResponse.User{}, err
	}
	resp := userResponse.FromDB(user)
	return resp, nil
}
