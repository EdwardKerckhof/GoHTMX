package auth

import (
	"github.com/EdwardKerckhof/gohtmx/internal/db"
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

// TODO: use service instead of store
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
}

func (h *authHandler) Login(ctx *gin.Context) {
}
