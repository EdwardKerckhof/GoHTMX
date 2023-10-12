package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
	"github.com/EdwardKerckhof/gohtmx/internal/handler/auth"
	todoHandler "github.com/EdwardKerckhof/gohtmx/internal/handler/todo"
	userHandler "github.com/EdwardKerckhof/gohtmx/internal/handler/user"
	authService "github.com/EdwardKerckhof/gohtmx/internal/service/auth"
	todoService "github.com/EdwardKerckhof/gohtmx/internal/service/todo"
	userService "github.com/EdwardKerckhof/gohtmx/internal/service/user"
)

const (
	apiPath = "/api/v1"
)

func New(store db.Store) *gin.Engine {
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

	// Setup services
	authService := authService.New(store)
	userService := userService.New(store)
	todoService := todoService.New(store)

	// Setup handlers
	authHandler := auth.New(apiRouter, authService)
	authHandler.RegisterRoutes()

	userHandler := userHandler.New(apiRouter, userService)
	userHandler.RegisterRoutes()

	todoHandler := todoHandler.New(apiRouter, todoService)
	todoHandler.RegisterRoutes()

	return router
}
