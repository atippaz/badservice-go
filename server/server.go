package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"bad-service-go/config"
	"bad-service-go/databases"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type echoServer struct {
	app  *echo.Echo
	db   databases.Database
	conf *config.Config
}

var (
	once   sync.Once
	server *echoServer
)

func NewEchoServer(conf *config.Config, db databases.Database) *echoServer {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)
	once.Do(func() {
		server = &echoServer{
			app:  echoApp,
			db:   db,
			conf: conf,
		}
	})
	return server
}
func (s *echoServer) Start() {
	qCh := make(chan os.Signal, 1)
	signal.Notify(
		qCh,
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	go s.GraceFullyShutdown(qCh)

	corsMiddleware := getCORSMiddleware(s.conf.Server.AllowOrigins)
	bodyMiddleware := getBodyLimitMiddleWare(s.conf.Server.BodyLimit)
	timeoutMiddleware := getTimeoutMiddleware(s.conf.Server.Timeout)
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())
	s.app.Use(bodyMiddleware)
	s.app.Use(corsMiddleware)
	s.app.Use(timeoutMiddleware)
	s.app.GET("/health", s.healthCheck)
	s.InitMatchRouter()
	s.InitMatchSetRouter()
	s.InittRoomRouter()

	s.httpListening()
}
func (s *echoServer) GraceFullyShutdown(quitCh chan os.Signal) {
	ctx := context.Background()
	<-quitCh
	s.app.Logger.Info("shutdown")
	if err := s.app.Shutdown(ctx); err != nil {
		s.app.Logger.Fatalf("Error: %s", err.Error())
	}
}
func (s *echoServer) httpListening() {
	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
	if err := s.app.Start(serverUrl); err != nil && err != http.ErrServerClosed {
		s.app.Logger.Fatalf("Error :%s", err.Error())

	}
}
func (s *echoServer) healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Ok is Good")
}
func getLoggerMiddleware() echo.MiddlewareFunc {
	return middleware.Logger()
}
func getTimeoutMiddleware(timeout time.Duration) echo.MiddlewareFunc {
	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Error Request Timeout",
		Timeout:      timeout * time.Second,
	})
}
func getCORSMiddleware(allowOrigins []string) echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: allowOrigins,
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	})
}

func getBodyLimitMiddleWare(bodyLimit string) echo.MiddlewareFunc {
	return middleware.BodyLimit(bodyLimit)
}
