package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	authRequest "github.com/EdwardKerckhof/gohtmx/internal/dto/request/auth"
	authService "github.com/EdwardKerckhof/gohtmx/internal/service/auth"
	"github.com/EdwardKerckhof/gohtmx/pkg/response"
)

const (
	prefix = "/auth"
)

type handler interface {
	RegisterRoutes()
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type authHandler struct {
	apiRouter *gin.RouterGroup
	service   authService.Service
}

func New(apiRouter *gin.RouterGroup, authService authService.Service) handler {
	return &authHandler{
		apiRouter: apiRouter,
		service:   authService,
	}
}

func (h *authHandler) RegisterRoutes() {
	// Rest API /api/v1/auth
	authRouter := h.apiRouter.Group(prefix)
	authRouter.POST("/register", h.Register)
	authRouter.POST("/login", h.Login)
}

func (h *authHandler) Register(ctx *gin.Context) {
	var req authRequest.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	user, err := h.service.Register(ctx, req)
	if err != nil {
		// TODO: better error handling
		ctx.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Success(user))
}

func (h *authHandler) Login(ctx *gin.Context) {
	panic("unimplemented")
}
