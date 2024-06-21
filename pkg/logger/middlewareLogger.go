package logger

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func MiddlewareLogger() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)

			stop := time.Now()

			req := c.Request()
			res := c.Response()

			logEvent := log.Info().
				// Str("status", "ok").
				Str("time_human", stop.Format(time.RFC3339Nano)).
				Str("path", req.URL.Path).
				Str("method", req.Method).
				Str("latency", strconv.FormatInt(int64(stop.Sub(start)), 10)).
				Str("latency_human", stop.Sub(start).String()).
				Int("response_code", res.Status).
				Int64("response_size", res.Size).
				Str("response_type", res.Header().Get("Content-Type")).
				Interface("response_headers", res.Header()).
				Interface("request_headers", req.Header).
				Interface("params", c.ParamNames()).
				Interface("query", req.URL.Query()).
				Interface("cookies", req.Cookies()).
				Interface("error", err).
				Str("ip", c.RealIP()).
				// Str("user", c.Get("user").(string)).
				Str("protocol", req.Proto).
				Str("host", req.Host)
				// Str("port", req.URL.Port())

			if req.TLS != nil {
				logEvent = logEvent.Uint16("tlsVersion", req.TLS.Version)
			}

			logEvent.Msg("request")

			return err
		}
	}
}
