package module

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/EdwardKerckhof/gohtmx/config"
	"github.com/EdwardKerckhof/gohtmx/internal/db"
	"github.com/EdwardKerckhof/gohtmx/internal/module/auth"
	"github.com/EdwardKerckhof/gohtmx/internal/module/todo"
	"github.com/EdwardKerckhof/gohtmx/internal/module/user"
	"github.com/EdwardKerckhof/gohtmx/pkg/token/paseto"
)

func InitModules(config config.Config, store db.Store, apiRouter *gin.RouterGroup) error {
	tokenMaker, err := paseto.NewMaker(config.Auth.TokenSymmetricKey)
	if err != nil {
		return fmt.Errorf("error creating token maker: %w", err)
	}

	auth.InitModule(config, store, apiRouter, tokenMaker)
	user.InitModule(store, apiRouter, tokenMaker)
	todo.InitModule(store, apiRouter, tokenMaker)

	return nil
}
