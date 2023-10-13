package user

import (
	"context"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
	userRequest "github.com/EdwardKerckhof/gohtmx/internal/dto/request/user"
	userModel "github.com/EdwardKerckhof/gohtmx/internal/model/user"
	"github.com/EdwardKerckhof/gohtmx/pkg/request"
)

type Service interface {
	Count(ctx context.Context) (int64, error)
	FindAll(ctx context.Context, req userRequest.FindAllRequest) ([]userModel.User, error)
	FindById(ctx context.Context, id request.IDRequest) (userModel.User, error)
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

func (s userService) FindAll(ctx context.Context, req userRequest.FindAllRequest) ([]userModel.User, error) {
	arg := db.FindAllUsersParams{
		Limit:  req.Size,
		Offset: (req.Page - 1) * req.Size,
	}

	users, err := s.store.FindAllUsers(ctx, arg)
	if err != nil {
		return []userModel.User{}, err
	}
	return userModel.FromDBList(users), nil
}

func (s userService) FindById(ctx context.Context, req request.IDRequest) (userModel.User, error) {
	id, err := req.ParseID()
	if err != nil {
		return userModel.User{}, err
	}

	user, err := s.store.FindUserById(ctx, id)
	if err != nil {
		return userModel.User{}, err
	}
	return userModel.FromDB(user), nil
}
