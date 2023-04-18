package config

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetRoute(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Static("/assets", "assets")
	e.File("/favicon.ico", "assets/favicon.svg")
}
