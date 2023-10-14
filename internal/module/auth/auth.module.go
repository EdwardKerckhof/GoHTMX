package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/EdwardKerckhof/gohtmx/config"
	"github.com/EdwardKerckhof/gohtmx/internal/db"
	"github.com/EdwardKerckhof/gohtmx/internal/module/auth/handler"
	"github.com/EdwardKerckhof/gohtmx/internal/module/auth/service"
	"github.com/EdwardKerckhof/gohtmx/pkg/token"
)

func InitModule(config config.Config, store db.Store, apiRouter *gin.RouterGroup, tokenMaker token.Maker) {
	service := service.New(config, store, tokenMaker)
	handler := handler.New(service, apiRouter)
	handler.RegisterRoutes()
}
