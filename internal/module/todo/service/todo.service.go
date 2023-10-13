package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
	"github.com/EdwardKerckhof/gohtmx/internal/module/todo/dto"
	"github.com/EdwardKerckhof/gohtmx/pkg/request"
)

type Service interface {
	Count(ctx context.Context) (int64, error)
	FindAll(ctx context.Context, req dto.FindAllRequest) ([]dto.Response, error)
	FindAllWithCount(ctx context.Context, req dto.FindAllRequest) ([]dto.Response, int64, error)
	FindById(ctx context.Context, req request.IDRequest) (dto.Response, error)
	Create(ctx context.Context, req dto.CreateRequest) (dto.Response, error)
	Update(ctx context.Context, idReq request.IDRequest, req dto.UpdateRequest) (dto.Response, error)
	Delete(ctx context.Context, req request.IDRequest) error
}

type todoService struct {
	store db.Store
}

func New(store db.Store) Service {
	return todoService{
		store: store,
	}
}

func (s todoService) Count(ctx context.Context) (int64, error) {
	return s.store.CountTodos(ctx)
}

func (s todoService) FindAll(ctx context.Context, req dto.FindAllRequest) ([]dto.Response, error) {
	arg := db.FindAllTodosParams{
		Limit:  req.Size,
		Offset: (req.Page - 1) * req.Size,
	}

	todos, err := s.store.FindAllTodos(ctx, arg)
	if err != nil {
		return []dto.Response{}, err
	}
	return dto.FromDBList(todos), nil
}

func (s todoService) FindAllWithCount(ctx context.Context, req dto.FindAllRequest) ([]dto.Response, int64, error) {
	todos, err := s.FindAll(ctx, req)
	if err != nil {
		return []dto.Response{}, 0, err
	}

	count, err := s.Count(ctx)
	if err != nil {
		return []dto.Response{}, 0, err
	}
	return todos, count, nil
}
func (s todoService) FindById(ctx context.Context, req request.IDRequest) (dto.Response, error) {
	id, err := req.ParseID()
	if err != nil {
		return dto.Response{}, err
	}

	todo, err := s.store.FindTodoById(ctx, id)
	if err != nil {
		return dto.Response{}, err
	}
	return dto.FromDB(todo), nil
}

func (s todoService) Create(ctx context.Context, req dto.CreateRequest) (dto.Response, error) {
	arg := db.CreateTodoParams{
		Title:  req.Title,
		UserID: uuid.New(), // TODO: Get user from context
	}

	todo, err := s.store.CreateTodo(ctx, arg)
	if err != nil {
		return dto.Response{}, err
	}
	return dto.FromDB(todo), nil
}

func (s todoService) Update(ctx context.Context, idReq request.IDRequest, req dto.UpdateRequest) (dto.Response, error) {
	id, err := idReq.ParseID()
	if err != nil {
		return dto.Response{}, err
	}

	arg := db.UpdateTodoParams{
		ID:        id,
		Title:     req.Title,
		Completed: req.Completed,
	}
	todo, err := s.store.UpdateTodo(ctx, arg)
	if err != nil {
		return dto.Response{}, err
	}
	return dto.FromDB(todo), nil
}

func (s todoService) Delete(ctx context.Context, req request.IDRequest) error {
	id, err := req.ParseID()
	if err != nil {
		return err
	}
	if err := s.store.DeleteTodo(ctx, id); err != nil {
		return err
	}
	return nil
}
