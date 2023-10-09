package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EdwardKerckhof/gohtmx/pkg/logger"
)

const (
	basePath = "/api/v1"
)

type Router interface {
	Engine() *gin.Engine
	BaseRouter() *gin.RouterGroup
}

type RouterImpl struct {
	engine *gin.Engine
	rg     *gin.RouterGroup
}

func New(logger logger.Logger) Router {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":      "healthy",
			"responseCode": http.StatusOK,
		})
	})

	return &RouterImpl{
		engine: router,
		rg:     router.Group(basePath),
	}
}

func (r *RouterImpl) Engine() *gin.Engine {
	return r.engine
}

func (r *RouterImpl) BaseRouter() *gin.RouterGroup {
	return r.rg
}
