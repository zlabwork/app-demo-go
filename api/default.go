package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"time"
)

func Home(c echo.Context) error {
	return c.String(http.StatusOK, os.Getenv("APP_NAME")+" is ok")
}

func SampleTemplate(c echo.Context) error {
	type wrap struct {
		Title    string
		CreateAt time.Time
	}
	data := &wrap{Title: "this is title", CreateAt: time.Now()}
	return c.Render(http.StatusOK, "layouts/main", data)
}
