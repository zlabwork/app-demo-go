package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func Logger() echo.MiddlewareFunc {
	f, _ := os.OpenFile("system.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: f,
	})
}
