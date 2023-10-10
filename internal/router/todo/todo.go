package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EdwardKerckhof/gohtmx/internal/domain"
	"github.com/EdwardKerckhof/gohtmx/internal/domain/todo"
	"github.com/EdwardKerckhof/gohtmx/internal/ports"
	"github.com/EdwardKerckhof/gohtmx/internal/postgres"
)

const (
	prefix = "/todos"
)

type handlerImpl struct {
	router *gin.RouterGroup
	store  *postgres.Store
}

func New(router *gin.RouterGroup, store *postgres.Store) ports.TodoRouter {
	return &handlerImpl{
		router: router,
		store:  store,
	}
}

func (c *handlerImpl) RegisterRoutes() {
	todoRouter := c.router.Group(prefix)
	todoRouter.GET("", c.FindAll)
	todoRouter.GET("/:id", c.FindById)
	todoRouter.POST("", c.Create)
	todoRouter.PUT("/:id", c.Update)
	todoRouter.DELETE("/:id", c.Delete)
}

func (c *handlerImpl) FindAll(ctx *gin.Context) {
	todos, err := c.store.TodoStore.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":      err.Error(),
			"responseCode": http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":      "success",
		"responseCode": http.StatusOK,
		"data":         todos,
	})
}

func (c *handlerImpl) FindById(ctx *gin.Context) {
	id, err := domain.ParseID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":      err.Error(),
			"responseCode": http.StatusBadRequest,
		})
		return
	}

	todo, err := c.store.TodoStore.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":      err.Error(),
			"responseCode": http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":      "success",
		"responseCode": http.StatusOK,
		"data":         todo,
	})
}

func (c *handlerImpl) Create(ctx *gin.Context) {
	var todo todo.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":      err.Error(),
			"responseCode": http.StatusBadRequest,
		})
		return
	}

	todo.GenerateID()

	if err := c.store.TodoStore.Create(&todo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":      err.Error(),
			"responseCode": http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":      "success",
		"responseCode": http.StatusOK,
		"data":         todo,
	})
}

func (c *handlerImpl) Update(ctx *gin.Context) {
	id, err := domain.ParseID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":      err.Error(),
			"responseCode": http.StatusBadRequest,
		})
		return
	}

	var todo todo.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":      err.Error(),
			"responseCode": http.StatusBadRequest,
		})
		return
	}

	todo.ID = id

	if err := c.store.TodoStore.Update(&todo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":      err.Error(),
			"responseCode": http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":      "success",
		"responseCode": http.StatusOK,
		"data":         todo,
	})
}

func (c *handlerImpl) Delete(ctx *gin.Context) {
	id, err := domain.ParseID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":      err.Error(),
			"responseCode": http.StatusBadRequest,
		})
		return
	}

	if err := c.store.TodoStore.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":      err.Error(),
			"responseCode": http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":      "success",
		"responseCode": http.StatusNoContent,
	})
}
