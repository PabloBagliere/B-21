package server

import (
	"github.com/PabloBagliere/B-21/pkg/logger"
	"github.com/PabloBagliere/B-21/pkg/router"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewServer(serviceName string) *echo.Echo {
	e := echo.New()
	e.Use(echoprometheus.NewMiddleware(serviceName))
	e.Use(logger.MiddlewareLogger())
	e.GET("/healthcheck", router.HealthCheck)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/metrics", echoprometheus.NewHandler())

	return e
}
