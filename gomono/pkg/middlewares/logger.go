package middlewares

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func makeLogEntry(c echo.Context) *log.Entry {
	if c == nil {
		return log.WithFields(log.Fields{})
	}

	return log.WithFields(log.Fields{
		"id":         c.Request().Header.Get(echo.HeaderXRequestID),
		"method":     c.Request().Method,
		"host":       c.Request().Host,
		"user_agent": c.Request().Header.Get("User-Agent"),
		"path":       c.Request().URL.Path,
		"remote_ip":  c.RealIP(),
	})
}

func requestLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		makeLogEntry(c).Info("REQ")
		return next(c)
	}
}

func RequestLogger() echo.MiddlewareFunc {
	return requestLogger
}
