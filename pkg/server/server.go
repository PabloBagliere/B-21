package server

import (
	"fmt"

	"github.com/PabloBagliere/B-21/pkg/config"
	"github.com/PabloBagliere/B-21/pkg/logger"
	"github.com/PabloBagliere/B-21/pkg/router"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewServer(serviceName string, logLevel zerolog.Level) *echo.Echo {
	e := echo.New()
	e.Use(echoprometheus.NewMiddleware(serviceName))
	e.Use(logger.MiddlewareLogger(serviceName, logLevel))
	e.GET("/healthcheck", router.HealthCheck)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/metrics", echoprometheus.NewHandler())

	p, err := config.GetConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}
	fmt.Println(p)

	return e
}
