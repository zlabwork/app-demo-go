package main

import (
	"app/config"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	e := echo.New()
	config.SetRoute(e)
	port := os.Getenv("APP_PORT")
	e.Logger.Fatal(e.Start(":" + port))
}
