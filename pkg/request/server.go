package request

import (
	"errors"
	"fmt"
	"github.com/adnanahmady/go-websocket-chat/config"
	"github.com/adnanahmady/go-websocket-chat/pkg/applog"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Router interface {
	GetEngine() *gin.Engine
}

var _ Router = (*Server)(nil)

type Server struct {
	cfg    *config.Config
	logger applog.Logger
	server *http.Server
	engine *gin.Engine
}

func NewServer(
	cfg *config.Config,
	logger applog.Logger,
) *Server {
	return &Server{
		cfg:    cfg,
		logger: logger,
		engine: gin.Default(),
		server: &http.Server{},
	}
}

func (s *Server) GetEngine() *gin.Engine {
	return s.engine
}

func (s *Server) prepare() {
	s.engine.Use(gin.Recovery())
	s.engine.Use(logMiddleware(s.logger))
}

func (s *Server) Start() error {
	host := fmt.Sprintf("%s:%d", s.cfg.App.Host, s.cfg.App.Port)
	s.server.Handler = s.engine
	s.server.Addr = host
	s.logger.Info("starting server on (%s)", host)
	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.logger.Error("failed to run server", "error", err)
		return err
	}
	return nil
}

func (s *Server) Shutdown() error {
	ctx, cancel := NewWithTimeout(5 * time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Error("failed to shutdown server", "error", err)
		return err
	}
	return nil
}
