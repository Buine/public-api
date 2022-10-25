package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tesis/user-ms/config"
	"github.com/tesis/user-ms/errors"
	"github.com/tesis/user-ms/routers"
	"gorm.io/gorm"
	"net/http"
)

type EchoServer struct {
	Server   *echo.Echo
	Database *gorm.DB
}

func (current *EchoServer) RunServer() {
	runServer := current.initServer()
	runServer()
}

func (current *EchoServer) RunServerAsync(serverOn chan bool) {
	startServer := current.initServer()
	go startServer()
	serverOn <- true
}

func (current *EchoServer) initServer() func() {
	// Server
	current.Server = echo.New()
	serverConfig := new(config.ServerConfig).Init()
	current.Server.HTTPErrorHandler = errors.CustomHTTPErrorHandler
	current.Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		ExposeHeaders: []string{"Authorization"},
	}))
	routers.ConfigRouters(current.Server)
	return func() {
		current.Server.Logger.Fatal(current.Server.Start(":" + serverConfig.Port))
	}
}
