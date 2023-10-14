package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EdwardKerckhof/gohtmx/internal/module/auth/dto"
	"github.com/EdwardKerckhof/gohtmx/internal/module/auth/service"
	"github.com/EdwardKerckhof/gohtmx/pkg/response"
)

const (
	prefix = "/auth"
)

type Handler interface {
	RegisterRoutes()
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	RefreshAccessToken(ctx *gin.Context)
}

type authHandler struct {
	service   service.Service
	apiRouter *gin.RouterGroup
}

func New(authService service.Service, apiRouter *gin.RouterGroup) Handler {
	return &authHandler{
		service:   authService,
		apiRouter: apiRouter,
	}
}

func (h *authHandler) RegisterRoutes() {
	// Rest API /api/v1/auth
	authRouter := h.apiRouter.Group(prefix)
	authRouter.POST("/register", h.Register)
	authRouter.POST("/login", h.Login)
	authRouter.POST("/refresh", h.RefreshAccessToken)
}

func (h *authHandler) Register(ctx *gin.Context) {
	var req dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	resp, err := h.service.Register(ctx, req)
	if err != nil {
		// TODO: better error handling
		ctx.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Success(resp))
}

func (h *authHandler) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	resp, err := h.service.Login(ctx, req)
	if err != nil {
		// TODO: better error handling
		ctx.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Success(resp))
}

func (h *authHandler) RefreshAccessToken(ctx *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	resp, err := h.service.RefreshAccessToken(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Success(resp))
}
