package main

import (
	"app/config"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {

	// Init
	e := echo.New()

	// Route
	config.SetRoute(e)

	// Start
	port := os.Getenv("APP_PORT")
	e.Logger.Fatal(e.Start(":" + port))
}
