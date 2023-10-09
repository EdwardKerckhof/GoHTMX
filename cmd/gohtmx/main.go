package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/EdwardKerckhof/gohtmx/config"
	"github.com/EdwardKerckhof/gohtmx/internal/postgres"
	"github.com/EdwardKerckhof/gohtmx/internal/router"
	"github.com/EdwardKerckhof/gohtmx/internal/server"
	"github.com/EdwardKerckhof/gohtmx/pkg/logger"
)

func main() {
	// Load config
	config, err := config.Load(".")
	if err != nil {
		log.Fatalf("error loading config: %s", err.Error())
	}

	// Setup app logger
	appLogger := logger.New(config)
	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s", config.Api.Version, config.Logger.Level, config.Api.Mode)

	// Create a new appRouter router
	appRouter := router.New(appLogger)

	// Create a new database connection
	store, err := postgres.NewStore(config)
	if err != nil {
		appLogger.Fatalf("error creating database connection: %s", err.Error())
	}

	// Setup handlers
	todoRouter := router.NewTodoRouter(appRouter, *store)
	todoRouter.RegisterRoutes()

	// Create a new server instance
	server := server.New(appRouter, config, appLogger)

	// Start the server
	server.Start()
	defer server.Stop()

	// Listen for OS signals to perform a graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	<-c
	log.Println("graceful shutdown...")
}
