package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
	"github.com/EdwardKerckhof/gohtmx/internal/module/auth/handler"
	"github.com/EdwardKerckhof/gohtmx/internal/module/auth/service"
)

func InitModule(store db.Store, apiRouter *gin.RouterGroup) {
	service := service.New(store)
	handler := handler.New(service, apiRouter)
	handler.RegisterRoutes()
}
