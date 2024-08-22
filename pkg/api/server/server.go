package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	oapimiddleware "github.com/oapi-codegen/echo-middleware"
	slogecho "github.com/samber/slog-echo"

	"github.com/nint8835/monticola/pkg/api/models"
	"github.com/nint8835/monticola/pkg/config"
)

var _ ServerInterface = (*Server)(nil)

type Server struct {
	config   *config.ServerConfig
	echoInst *echo.Echo
}

func (s *Server) Start() error {
	return s.echoInst.Start(s.config.ListenAddress)
}

func (s *Server) GetTest(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.Test{Message: "Hello, world!"})
}

func New(c *config.ServerConfig) (*Server, error) {
	echoInst := echo.New()
	serverInst := &Server{
		config:   c,
		echoInst: echoInst,
	}

	echoInst.HideBanner = true

	echoInst.Use(slogecho.New(slog.Default()))

	oapi, err := GetSwagger()
	if err != nil {
		return nil, fmt.Errorf("error getting openapi spec: %w", err)
	}

	oapi.Servers = nil

	echoInst.Use(oapimiddleware.OapiRequestValidator(oapi))

	RegisterHandlers(echoInst, serverInst)

	return serverInst, nil
}
