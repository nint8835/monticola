package server

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	slogecho "github.com/samber/slog-echo"

	"github.com/nint8835/monticola/pkg/config"
)

type Server struct {
	config   *config.ServerConfig
	echoInst *echo.Echo
}

func (s *Server) Start() error {
	return s.echoInst.Start(s.config.ListenAddress)
}

func New(c *config.ServerConfig) (*Server, error) {
	echoInst := echo.New()
	serverInst := &Server{
		config:   c,
		echoInst: echoInst,
	}

	echoInst.HideBanner = true

	echoInst.Use(slogecho.New(slog.Default()))

	return serverInst, nil
}
