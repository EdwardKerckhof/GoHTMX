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
	// Update(ctx *gin.Context)
	// Delete(ctx *gin.Context)
}

type handlerImpl struct {
	router *gin.RouterGroup
	store  *db.Store
}

func New(router *gin.RouterGroup, store *db.Store) handler {
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
	// todoRouter.PUT("/:id", c.Update)
	// todoRouter.DELETE("/:id", c.Delete)
}

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
	var req findByIdRequest
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

// func (c *handlerImpl) Update(ctx *gin.Context) {
// 	id, err := domain.ParseID(ctx.Param("id"))
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"message":      err.Error(),
// 			"responseCode": http.StatusBadRequest,
// 		})
// 		return
// 	}

// 	var todo todo.Todo
// 	if err := ctx.ShouldBindJSON(&todo); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"message":      err.Error(),
// 			"responseCode": http.StatusBadRequest,
// 		})
// 		return
// 	}

// 	todo.ID = id

// 	if err := c.store.TodoStore.Update(&todo); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"message":      err.Error(),
// 			"responseCode": http.StatusInternalServerError,
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"message":      "success",
// 		"responseCode": http.StatusOK,
// 		"data":         todo,
// 	})
// }

// func (c *handlerImpl) Delete(ctx *gin.Context) {
// 	id, err := domain.ParseID(ctx.Param("id"))
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"message":      err.Error(),
// 			"responseCode": http.StatusBadRequest,
// 		})
// 		return
// 	}

// 	if err := c.store.TodoStore.Delete(id); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"message":      err.Error(),
// 			"responseCode": http.StatusInternalServerError,
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"message":      "success",
// 		"responseCode": http.StatusNoContent,
// 	})
// }
