package todo

import (
	"github.com/gin-gonic/gin"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
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
