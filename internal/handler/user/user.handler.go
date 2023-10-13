package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	userReq "github.com/EdwardKerckhof/gohtmx/internal/dto/request/user"
	userService "github.com/EdwardKerckhof/gohtmx/internal/service/user"
	"github.com/EdwardKerckhof/gohtmx/pkg/request"
	"github.com/EdwardKerckhof/gohtmx/pkg/response"
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
	service   userService.Service
}

func New(apiRouter *gin.RouterGroup, service userService.Service) handler {
	return &userHandler{
		apiRouter: apiRouter,
		service:   service,
	}
}

func (h *userHandler) RegisterRoutes() {
	// Rest API /api/v1/users
	userRouter := h.apiRouter.Group(prefix)
	userRouter.GET("", h.FindAll)
	userRouter.GET("/:id", h.FindById)
}

func (h *userHandler) FindAll(ctx *gin.Context) {
	var req userReq.FindAllRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	users, err := h.service.FindAll(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}

	usersCount, err := h.service.Count(ctx)
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

	user, err := h.service.FindById(ctx, req)
	if err != nil {
		// TODO: better error handling
		ctx.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Success(user))
}
