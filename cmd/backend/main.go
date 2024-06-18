package main

import (
	"fmt"

	"github.com/PabloBagliere/B-21/internal/router"
	"github.com/PabloBagliere/B-21/pkg/config"
	"github.com/PabloBagliere/B-21/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func main() {
	// ...
	// HealthCheck(e)
	// ...
	e := echo.New()
	e.Use(logger.MiddlewareLogger("MyServiceName", zerolog.InfoLevel))
	router.HealthCheck(e)
	p, err := config.GetConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}
	fmt.Println(p)
	e.Logger.Fatal(e.Start(":8080"))

}
