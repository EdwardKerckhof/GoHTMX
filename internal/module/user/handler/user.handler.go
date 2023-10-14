package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EdwardKerckhof/gohtmx/internal/middleware"
	"github.com/EdwardKerckhof/gohtmx/internal/module/user/dto"
	"github.com/EdwardKerckhof/gohtmx/internal/module/user/service"
	"github.com/EdwardKerckhof/gohtmx/pkg/request"
	"github.com/EdwardKerckhof/gohtmx/pkg/response"
	"github.com/EdwardKerckhof/gohtmx/pkg/token"
)

const (
	prefix = "/users"
)

type Handler interface {
	RegisterRoutes()
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
}

type userHandler struct {
	service    service.Service
	apiRouter  *gin.RouterGroup
	tokenMaker token.Maker
}

func New(service service.Service, apiRouter *gin.RouterGroup, tokenMaker token.Maker) Handler {
	return &userHandler{
		service:    service,
		apiRouter:  apiRouter,
		tokenMaker: tokenMaker,
	}
}

func (h *userHandler) RegisterRoutes() {
	// Rest API /api/v1/users
	userRouter := h.apiRouter.Group(prefix)
	userRouter.Use(middleware.AuthMiddleware(h.tokenMaker)).GET("", h.FindAll)
	userRouter.GET("/:id", h.FindById)
}

func (h *userHandler) FindAll(ctx *gin.Context) {
	var req dto.FindAllRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	users, count, err := h.service.FindAllWithCount(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Paginated(users, count, req.PaginationRequest))
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
