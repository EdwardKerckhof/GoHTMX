package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
	"github.com/EdwardKerckhof/gohtmx/pkg/request"
	"github.com/EdwardKerckhof/gohtmx/pkg/response"
)

const (
	prefix = "/todos"
)

type handler interface {
	RegisterRoutes()
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

// TODO: use service instead of store
type todoHandler struct {
	apiRouter *gin.RouterGroup
	store     *db.Store
}

func New(apiRouter *gin.RouterGroup, store *db.Store) handler {
	return &todoHandler{
		apiRouter: apiRouter,
		store:     store,
	}
}

func (h *todoHandler) RegisterRoutes() {
	// Rest API /api/v1/todos
	todoRouter := h.apiRouter.Group(prefix)
	todoRouter.GET("", h.FindAll)
	todoRouter.GET("/:id", h.FindById)
	todoRouter.POST("", h.Create)
	todoRouter.PUT("/:id", h.Update)
	todoRouter.DELETE("/:id", h.Delete)
}

func (c *todoHandler) FindAll(ctx *gin.Context) {
	var req findAllRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	arg := db.FindAllTodosParams{
		Limit:  req.Size,
		Offset: (req.Page - 1) * req.Size,
	}
	todos, err := c.store.FindAllTodos(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}

	todosCount, err := c.store.CountTodos(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Paginated(todos, todosCount, req.PaginationRequest))
}

func (c *todoHandler) FindById(ctx *gin.Context) {
	var req request.IDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	id, err := req.ParseID()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	todo, err := c.store.FindTodoById(ctx, id)
	if err != nil {
		// TODO: better error handling
		ctx.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Success(todo))
}

func (c *todoHandler) Create(ctx *gin.Context) {
	var req createRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	arg := db.CreateTodoParams{
		Title:  req.Title,
		UserID: uuid.New(), // TODO: Get user from context
	}

	todo, err := c.store.CreateTodo(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Success(todo))
}

func (c *todoHandler) Update(ctx *gin.Context) {
	var req updateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}
	var idReq request.IDRequest
	if err := ctx.ShouldBindUri(&idReq); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	id, err := idReq.ParseID()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	arg := db.UpdateTodoParams{
		ID:        id,
		Title:     req.Title,
		Completed: req.Completed,
	}
	todo, err := c.store.UpdateTodo(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Success(todo))
}

func (c *todoHandler) Delete(ctx *gin.Context) {
	var req request.IDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	id, err := req.ParseID()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	if err := c.store.DeleteTodo(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
