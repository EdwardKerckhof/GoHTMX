package module

import (
	"github.com/gin-gonic/gin"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
	"github.com/EdwardKerckhof/gohtmx/internal/module/auth"
	"github.com/EdwardKerckhof/gohtmx/internal/module/todo"
	"github.com/EdwardKerckhof/gohtmx/internal/module/user"
)

func InitModules(store db.Store, apiRouter *gin.RouterGroup) {
	auth.InitModule(store, apiRouter)
	user.InitModule(store, apiRouter)
	todo.InitModule(store, apiRouter)
}
