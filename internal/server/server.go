package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/EdwardKerckhof/gohtmx/config"
	"github.com/EdwardKerckhof/gohtmx/internal/router"
	"github.com/EdwardKerckhof/gohtmx/pkg/logger"
)

type HttpServer interface {
	Start()
	Stop()
}

type httpServer struct {
	config *config.Config
	logger logger.Logger
	server *http.Server
}

func New(router router.Router, config *config.Config, logger logger.Logger) HttpServer {
	return &httpServer{
		config: config,
		logger: logger,
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", config.Api.Port),
			Handler: router.Engine(),
		},
	}
}

func (s *httpServer) Start() {
	port := s.config.Api.Port
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatalf(
				"failed to stater HttpServer listen port %d, err=%s",
				port, err.Error(),
			)
		}
	}()

	s.logger.Infof("server listening on port: %d", port)
}

func (s *httpServer) Stop() {
	ctx, cancel := context.WithTimeout(
		context.Background(), time.Duration(3)*time.Second,
	)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Fatalf("server shutdown failed, err=%s", err.Error())
	}
}
