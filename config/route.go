package config

import (
	"app/api"
	"app/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetRoute(e *echo.Echo) {

	// Middleware
	e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	r := e.Group("/api")
	r.Use(middleware.Jwt())
	r.GET("", api.Restricted)

	// defined
	e.POST("/token", api.Token)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// assets
	e.Static("/assets", "assets")
	e.File("/favicon.ico", "assets/favicon.svg")
}
