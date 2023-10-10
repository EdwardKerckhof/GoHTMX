package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/EdwardKerckhof/gohtmx/config"
	"github.com/EdwardKerckhof/gohtmx/pkg/logger"
	"github.com/gin-gonic/gin"
)

type HttpServer interface {
	Start()
	Stop()
}

type httpServer struct {
	logger logger.Logger
	server *http.Server
}

func New(router *gin.Engine, config *config.Config, logger logger.Logger) HttpServer {
	return &httpServer{
		logger: logger,
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", config.Api.Port),
			Handler: router,
		},
	}
}

func (s *httpServer) Start() {
	addr := s.server.Addr
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatalf(
				"failed to stater HttpServer listen port %s, err=%s",
				addr, err.Error(),
			)
		}
	}()

	s.logger.Infof("server listening on port: %s", addr)
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
