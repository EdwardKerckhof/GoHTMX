package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
	"github.com/EdwardKerckhof/gohtmx/internal/domain"
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
	RenderIndex(ctx *gin.Context)
}

type handlerImpl struct {
	apiRouter  *gin.RouterGroup
	viewRouter *gin.RouterGroup
	store      *db.Store
}

func New(apiRouter *gin.RouterGroup, viewRouter *gin.RouterGroup, store *db.Store) handler {
	return &handlerImpl{
		apiRouter:  apiRouter,
		viewRouter: viewRouter,
		store:      store,
	}
}

func (c *handlerImpl) RegisterRoutes() {
	// /api/v1/todos
	todoRouter := c.apiRouter.Group(prefix)
	todoRouter.GET("", c.FindAll)
	todoRouter.GET("/:id", c.FindById)
	todoRouter.POST("", c.Create)
	todoRouter.PUT("/:id", c.Update)
	todoRouter.DELETE("/:id", c.Delete)

	// /todos
	viewsRouter := c.viewRouter.Group(prefix)
	viewsRouter.GET("", c.RenderIndex)
}

// REST API
func (c *handlerImpl) FindAll(ctx *gin.Context) {
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

func (c *handlerImpl) FindById(ctx *gin.Context) {
	var req idRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	id, err := domain.ParseID(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	todo, err := c.store.FindTodoById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Success(todo))
}

func (c *handlerImpl) Create(ctx *gin.Context) {
	var req createRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	arg := db.CreateTodoParams{
		ID:    domain.GenerateID(),
		Title: req.Title,
	}
	todo, err := c.store.CreateTodo(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Success(todo))
}

func (c *handlerImpl) Update(ctx *gin.Context) {
	var req updateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}
	var idReq idRequest
	if err := ctx.ShouldBindUri(&idReq); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	id, err := domain.ParseID(idReq.ID)
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

func (c *handlerImpl) Delete(ctx *gin.Context) {
	var req idRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	id, err := domain.ParseID(req.ID)
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

// Views
func (c *handlerImpl) RenderIndex(ctx *gin.Context) {
	arg := db.FindAllTodosParams{
		Limit:  50,
		Offset: 0,
	}
	// TODO: use service
	todos, err := c.store.FindAllTodos(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}

	data := gin.H{
		"title": "Todo App",
		"todos": todos,
	}

	ctx.HTML(http.StatusOK, "todos/index.html", data)
}
