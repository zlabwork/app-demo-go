package main

import (
	"app/config"
	"github.com/labstack/echo/v4"
	"io"
	"os"
	"text/template"
)

type Template struct {
	templates *template.Template
}

func (tpl *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// @docs https://colobu.com/2016/10/09/Go-embedded-template-best-practices/

	// parse template every time
	if os.Getenv("APP_ENV") == "dev" {
		tpl.templates = template.Must(template.ParseGlob("templates/*/*.html"))
	}
	return tpl.templates.ExecuteTemplate(w, name, data)
}

func main() {

	// Init
	e := echo.New()

	// Template
	e.Renderer = &Template{templates: template.Must(template.ParseGlob("templates/*/*.html"))}

	// Route
	config.SetRoute(e)

	// Start
	port := os.Getenv("APP_PORT")
	e.Logger.Fatal(e.Start(":" + port))
}
