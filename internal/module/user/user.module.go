package user

import (
	"github.com/EdwardKerckhof/gohtmx/internal/db"
	"github.com/EdwardKerckhof/gohtmx/internal/module/user/handler"
	"github.com/EdwardKerckhof/gohtmx/internal/module/user/service"
	"github.com/gin-gonic/gin"
)

func InitModule(store db.Store, apiRouter *gin.RouterGroup) {
	service := service.New(store)
	handler := handler.New(service, apiRouter)
	handler.RegisterRoutes()
}
