package auth

import (
	"net/http"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
	"github.com/EdwardKerckhof/gohtmx/pkg/response"
	"github.com/gin-gonic/gin"
)

const (
	prefix = "/auth"
)

type handler interface {
	RegisterRoutes()
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

// TODO: use service instead of store and in the service use DTOs
type authHandler struct {
	apiRouter *gin.RouterGroup
	store     *db.Store
}

func New(apiRouter *gin.RouterGroup, store *db.Store) handler {
	return &authHandler{
		apiRouter: apiRouter,
		store:     store,
	}
}

func (h *authHandler) RegisterRoutes() {
	// Rest API /api/v1/auth
	authRouter := h.apiRouter.Group(prefix)
	authRouter.POST("/register", h.Register)
	authRouter.POST("/login", h.Login)
}

func (h *authHandler) Register(ctx *gin.Context) {
	var req registerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	hashedPassword, err := req.HashPassword()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}

	arg := db.CreateUserParams{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
	}
	user, err := h.store.CreateUser(ctx, arg)
	if err != nil {
		// TODO: better error handling
		ctx.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Success(user))
}

func (h *authHandler) Login(ctx *gin.Context) {
}
