package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EdwardKerckhof/gohtmx/internal/dto/request"
	todoRequest "github.com/EdwardKerckhof/gohtmx/internal/dto/request/todo"
	"github.com/EdwardKerckhof/gohtmx/internal/dto/response"
	todoService "github.com/EdwardKerckhof/gohtmx/internal/service/todo"
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
	service   todoService.Service
}

func New(apiRouter *gin.RouterGroup, service todoService.Service) handler {
	return &todoHandler{
		apiRouter: apiRouter,
		service:   service,
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
	var req todoRequest.FindAllRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	todos, err := c.service.FindAll(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}

	todosCount, err := c.service.Count(ctx)
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

	todo, err := c.service.FindById(ctx, req)
	if err != nil {
		// TODO: better error handling
		ctx.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Success(todo))
}

func (c *todoHandler) Create(ctx *gin.Context) {
	var req todoRequest.CreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	todo, err := c.service.Create(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Success(todo))
}

func (c *todoHandler) Update(ctx *gin.Context) {
	var req todoRequest.UpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}
	var idReq request.IDRequest
	if err := ctx.ShouldBindUri(&idReq); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	todo, err := c.service.Update(ctx, idReq, req)
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

	if err := c.service.Delete(ctx, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
