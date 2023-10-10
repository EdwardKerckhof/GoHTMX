package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
	"github.com/EdwardKerckhof/gohtmx/internal/handler/todo"
)

const (
	apiPath  = "/api/v1"
	viewPath = ""
)

func New(store *db.Store) *gin.Engine {
	router := gin.New()

	// Middleware
	router.Use(gin.Recovery(), gin.Logger())

	// Healthcheck
	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":      "healthy",
			"responseCode": http.StatusOK,
		})
	})

	// Setup template rendering
	router.LoadHTMLGlob("web/templates/**/*")

	// Setup base routers
	apiRouter := router.Group(apiPath)
	viewRouter := router.Group(viewPath)

	// Setup handlers
	todoHandler := todo.New(apiRouter, viewRouter, store)
	todoHandler.RegisterRoutes()

	return router
}
