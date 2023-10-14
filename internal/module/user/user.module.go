package user

import (
	"github.com/gin-gonic/gin"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
	"github.com/EdwardKerckhof/gohtmx/internal/module/user/handler"
	"github.com/EdwardKerckhof/gohtmx/internal/module/user/service"
	"github.com/EdwardKerckhof/gohtmx/pkg/token"
)

func InitModule(store db.Store, apiRouter *gin.RouterGroup, tokenMaker token.Maker) {
	service := service.New(store)
	handler := handler.New(service, apiRouter, tokenMaker)
	handler.RegisterRoutes()
}
