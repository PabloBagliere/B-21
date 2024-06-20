package main

import (
	"fmt"

	_ "github.com/PabloBagliere/B-21/api/base"
	"github.com/PabloBagliere/B-21/internal/router"
	"github.com/PabloBagliere/B-21/pkg/config"
	"github.com/PabloBagliere/B-21/pkg/logger"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger API Example
// @version 0.0.1
// @description

// @contact.name API Support
// @contact.url https://github.com/PabloBagliere/B-21/issues
// @contact.email pablo.bagliere2@gmail.com

// @license.name MIT
// @license.url https://github.com/PabloBagliere/B-21/licence

// @host localhost:8080
func main() {
	e := echo.New()
	e.Use(echoprometheus.NewMiddleware("MyServiceName"))
	e.Use(logger.MiddlewareLogger("MyServiceName", zerolog.InfoLevel))
	e.GET("/healthcheck", router.HealthCheck)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/metrics", echoprometheus.NewHandler())
	p, err := config.GetConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}
	fmt.Println(p)
	e.Logger.Fatal(e.Start(":8080"))

}
