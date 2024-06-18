package logger

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func MiddlewareLogger(serviceName string, logLevel zerolog.Level) echo.MiddlewareFunc {
	// Configura el logger global
	InitLogger(Config{
		ServiceName: serviceName,
		LogLevel:    logLevel,
	})

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

// return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
// 	LogURI:    true,
// 	LogStatus: true,
// 	LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
// 		logEvent := log.Info().
// 			// Str("status", "ok").
// 			Str("path", v.URI).
// 			Str("method", c.Request().Method).
// 			Int64("responseTime", v.Latency.Milliseconds()).
// 			Int("responseCode", c.Response().Status).
// 			Int64("responseSize", c.Response().Size).
// 			Str("responseType", c.Response().Header().Get("Content-Type")).
// 			Interface("responseHeaders", c.Response().Header()).
// 			Interface("requestHeaders", c.Request().Header).
// 			// Str("requestBody", v.RequestBody).
// 			Interface("params", c.ParamNames()).
// 			Interface("query", c.QueryParams()).
// 			Interface("cookies", c.Cookies()).
// 			Str("ip", c.RealIP()).
// 			// Str("user", v.UserID).
// 			Interface("error", v.Error).
// 			Str("protocol", c.Request().Proto).
// 			Str("host", c.Request().Host).
// 			// Str("port", c.Request().URL.Port()).
// 			Bool("tls", c.Request().TLS != nil)

// 		// if c.Request().TLS != nil {
// 		// 	logEvent = logEvent.Str("tlsVersion", c.Request().TLS.Version)
// 		// }

// 		logEvent.Msg("request")
// 		return nil
// 	},
// })
