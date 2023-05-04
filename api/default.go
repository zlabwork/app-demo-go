package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func Home(c echo.Context) error {
	return c.String(http.StatusOK, os.Getenv("APP_NAME")+" is ok")
}
