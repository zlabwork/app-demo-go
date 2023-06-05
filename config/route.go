package config

import (
	"app/api"
	"app/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func SetRoute(e *echo.Echo) {

	// Middleware
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Cors())
	e.Pre(echoMiddleware.RemoveTrailingSlash())

	// defined
	e.GET("/", api.Home)
	e.GET("/sample", api.SampleTemplate)
	e.GET("/public_key", api.PublicKey)
	e.POST("/token", api.Token)
	e.PUT("/token", api.RefreshToken)
	e.Static("/assets", "assets")
	e.File("/favicon.ico", "assets/favicon.svg")

	// e.Use(middleware.Recover())
	r := e.Group("/api")
	r.Use(middleware.Jwt())
	r.GET("/me", api.Me)
}
