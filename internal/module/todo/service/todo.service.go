package service

import (
	"github.com/gin-gonic/gin"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
	"github.com/EdwardKerckhof/gohtmx/internal/middleware"
	"github.com/EdwardKerckhof/gohtmx/internal/module/todo/dto"
	"github.com/EdwardKerckhof/gohtmx/pkg/request"
	"github.com/EdwardKerckhof/gohtmx/pkg/token"
)

type Service interface {
	Count(ctx *gin.Context) (int64, error)
	FindAll(ctx *gin.Context, req dto.FindAllRequest) ([]dto.Response, error)
	FindAllWithCount(ctx *gin.Context, req dto.FindAllRequest) ([]dto.Response, int64, error)
	FindById(ctx *gin.Context, req request.IDRequest) (dto.Response, error)
	Create(ctx *gin.Context, req dto.CreateRequest) (dto.Response, error)
	Update(ctx *gin.Context, idReq request.IDRequest, req dto.UpdateRequest) (dto.Response, error)
	Delete(ctx *gin.Context, req request.IDRequest) error
}

type todoService struct {
	store db.Store
}

func New(store db.Store) Service {
	return todoService{
		store: store,
	}
}

func (s todoService) Count(ctx *gin.Context) (int64, error) {
	authPayload := ctx.MustGet(middleware.PayloadKey).(*token.Payload)
	return s.store.CountTodos(ctx, authPayload.UserID)
}

func (s todoService) FindAll(ctx *gin.Context, req dto.FindAllRequest) ([]dto.Response, error) {
	authPayload := ctx.MustGet(middleware.PayloadKey).(*token.Payload)
	arg := db.FindAllTodosParams{
		UserID: authPayload.UserID,
		Limit:  req.Size,
		Offset: (req.Page - 1) * req.Size,
	}

	todos, err := s.store.FindAllTodos(ctx, arg)
	if err != nil {
		return []dto.Response{}, err
	}
	return dto.NewResponseList(todos), nil
}

func (s todoService) FindAllWithCount(ctx *gin.Context, req dto.FindAllRequest) ([]dto.Response, int64, error) {
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

func (s todoService) FindById(ctx *gin.Context, req request.IDRequest) (dto.Response, error) {
	id, err := req.ParseID()
	if err != nil {
		return dto.Response{}, err
	}

	todo, err := s.store.FindTodoById(ctx, id)
	if err != nil {
		return dto.Response{}, err
	}
	return dto.NewResponse(todo), nil
}

func (s todoService) Create(ctx *gin.Context, req dto.CreateRequest) (dto.Response, error) {
	authPayload := ctx.MustGet(middleware.PayloadKey).(*token.Payload)
	arg := db.CreateTodoParams{
		Title:  req.Title,
		UserID: authPayload.UserID,
	}

	todo, err := s.store.CreateTodo(ctx, arg)
	if err != nil {
		return dto.Response{}, err
	}
	return dto.NewResponse(todo), nil
}

func (s todoService) Update(ctx *gin.Context, idReq request.IDRequest, req dto.UpdateRequest) (dto.Response, error) {
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
	return dto.NewResponse(todo), nil
}

func (s todoService) Delete(ctx *gin.Context, req request.IDRequest) error {
	id, err := req.ParseID()
	if err != nil {
		return err
	}
	if err := s.store.DeleteTodo(ctx, id); err != nil {
		return err
	}
	return nil
}
