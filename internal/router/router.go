package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EdwardKerckhof/gohtmx/internal/postgres"
	"github.com/EdwardKerckhof/gohtmx/internal/router/todo"
)

const (
	basePath = "/api/v1"
)

func New(store *postgres.Store) *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":      "healthy",
			"responseCode": http.StatusOK,
		})
	})

	baseRouter := router.Group(basePath)

	// Setup handlers
	todoRouter := todo.New(baseRouter, store)
	todoRouter.RegisterRoutes()

	return router
}
