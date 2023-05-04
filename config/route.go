package config

import (
	"app/api"
	"app/middleware"
	"github.com/labstack/echo/v4"
)

func SetRoute(e *echo.Echo) {

	// Middleware
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Cors())
	// e.Use(middleware.Recover())
	r := e.Group("/api")
	r.Use(middleware.Jwt())
	r.GET("/me", api.Me)

	// defined
	e.GET("/", api.Home)
	e.POST("/token", api.Token)

	// assets
	e.Static("/assets", "assets")
	e.File("/favicon.ico", "assets/favicon.svg")
}
