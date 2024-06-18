package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type healthcheck struct {
	Status string `json:"status" xml:"status"`
}

func HealthCheck(e *echo.Echo) {
	e.GET("/healthcheck", func(c echo.Context) error {
		status := &healthcheck{Status: "ok"}
		return c.JSON(http.StatusOK, status)
	})
}
