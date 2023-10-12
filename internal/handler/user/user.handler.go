package user

import (
	"net/http"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
	"github.com/EdwardKerckhof/gohtmx/pkg/request"
	"github.com/EdwardKerckhof/gohtmx/pkg/response"
	"github.com/gin-gonic/gin"
)

const (
	prefix = "/users"
)

type handler interface {
	RegisterRoutes()
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
}

type userHandler struct {
	apiRouter *gin.RouterGroup
	store     *db.Store
}

func New(apiRouter *gin.RouterGroup, store *db.Store) handler {
	return &userHandler{
		apiRouter: apiRouter,
		store:     store,
	}
}

func (h *userHandler) RegisterRoutes() {
	// Rest API /api/v1/users
	userRouter := h.apiRouter.Group(prefix)
	userRouter.GET("", h.FindAll)
	userRouter.GET("/:id", h.FindById)
}

func (h *userHandler) FindAll(ctx *gin.Context) {
	var req findAllRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	arg := db.FindAllUsersParams{
		Limit:  req.Size,
		Offset: (req.Page - 1) * req.Size,
	}
	users, err := h.store.FindAllUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}

	usersCount, err := h.store.CountUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Paginated(users, usersCount, req.PaginationRequest))
}

func (h *userHandler) FindById(ctx *gin.Context) {
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

	user, err := h.store.FindUserById(ctx, id)
	if err != nil {
		// TODO: better error handling
		ctx.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Success(user))
}
