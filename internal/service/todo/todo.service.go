package todo

import (
	"context"

	"github.com/google/uuid"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
	todoRequest "github.com/EdwardKerckhof/gohtmx/internal/dto/request/todo"
	todoModel "github.com/EdwardKerckhof/gohtmx/internal/model/todo"
	"github.com/EdwardKerckhof/gohtmx/pkg/request"
)

type Service interface {
	Count(ctx context.Context) (int64, error)
	FindAll(ctx context.Context, req todoRequest.FindAllRequest) ([]todoModel.Todo, error)
	FindById(ctx context.Context, req request.IDRequest) (todoModel.Todo, error)
	Create(ctx context.Context, req todoRequest.CreateRequest) (todoModel.Todo, error)
	Update(ctx context.Context, idReq request.IDRequest, req todoRequest.UpdateRequest) (todoModel.Todo, error)
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

func (s todoService) FindAll(ctx context.Context, req todoRequest.FindAllRequest) ([]todoModel.Todo, error) {
	arg := db.FindAllTodosParams{
		Limit:  req.Size,
		Offset: (req.Page - 1) * req.Size,
	}

	todos, err := s.store.FindAllTodos(ctx, arg)
	if err != nil {
		return []todoModel.Todo{}, err
	}
	return todoModel.FromDBList(todos), nil
}

func (s todoService) FindById(ctx context.Context, req request.IDRequest) (todoModel.Todo, error) {
	id, err := req.ParseID()
	if err != nil {
		return todoModel.Todo{}, err
	}

	todo, err := s.store.FindTodoById(ctx, id)
	if err != nil {
		return todoModel.Todo{}, err
	}
	return todoModel.FromDB(todo), nil
}

func (s todoService) Create(ctx context.Context, req todoRequest.CreateRequest) (todoModel.Todo, error) {
	arg := db.CreateTodoParams{
		Title:  req.Title,
		UserID: uuid.New(), // TODO: Get user from context
	}

	todo, err := s.store.CreateTodo(ctx, arg)
	if err != nil {
		return todoModel.Todo{}, err
	}
	return todoModel.FromDB(todo), nil
}

func (s todoService) Update(ctx context.Context, idReq request.IDRequest, req todoRequest.UpdateRequest) (todoModel.Todo, error) {
	id, err := idReq.ParseID()
	if err != nil {
		return todoModel.Todo{}, err
	}

	arg := db.UpdateTodoParams{
		ID:        id,
		Title:     req.Title,
		Completed: req.Completed,
	}
	todo, err := s.store.UpdateTodo(ctx, arg)
	if err != nil {
		return todoModel.Todo{}, err
	}
	return todoModel.FromDB(todo), nil
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
