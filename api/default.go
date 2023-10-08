package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"runtime"
	"time"
)

func Home(c echo.Context) error {
	type wrap struct {
		GolangVersion string
		EchoVersion   string
	}

	data := &wrap{GolangVersion: runtime.Version(), EchoVersion: echo.Version}
	return c.Render(http.StatusOK, "default/example.html", data)
}

func SampleTemplate(c echo.Context) error {
	type wrap struct {
		Title    string
		CreateAt time.Time
	}
	data := &wrap{Title: "this is title", CreateAt: time.Now()}
	return c.Render(http.StatusOK, "default/main.html", data)
}
