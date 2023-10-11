package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
	"github.com/EdwardKerckhof/gohtmx/internal/handler/todo"
)

const (
	apiPath = "/api/v1"
)

func New(store *db.Store) *gin.Engine {
	router := gin.New()

	// Middleware
	router.Use(gin.Recovery(), gin.Logger())

	// Health check
	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":      "healthy",
			"responseCode": http.StatusOK,
		})
	})

	// Setup base routers
	apiRouter := router.Group(apiPath)

	// Setup handlers
	todoHandler := todo.New(apiRouter, store)
	todoHandler.RegisterRoutes()

	return router
}
