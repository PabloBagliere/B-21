package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type healthcheck struct {
	Status string `json:"status" xml:"status"`
}

// @Summary Show healthcheck status
// @Description get healthcheck status
// @ID get-healthcheck
// @Produce  json
// @Success 200 {object} healthcheck
// @Router /healthcheck [get]
func HealthCheck(c echo.Context) error {
	status := &healthcheck{Status: "ok"}
	return c.JSON(http.StatusOK, status)
}
