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
	logger := logger.New(config)
	logger.InitLogger()
	logger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s", config.Api.Version, config.Logger.Level, config.Api.Mode)

	// Create a new database connection
	store, err := postgres.NewStore(config)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	// Create a new router router
	router := router.New(store)

	// Create a new server instance
	server := server.New(router, config, logger)

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
