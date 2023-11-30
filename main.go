package main

import (
	"app/config"
	"github.com/labstack/echo/v4"
	"io"
	"os"
	"path/filepath"
	"text/template"
)

type Template struct {
	templates *template.Template
}

func (tpl *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// @docs https://colobu.com/2016/10/09/Go-embedded-template-best-practices/

	layouts, _ := filepath.Glob("templates/layouts/*.html")
	widgets, _ := filepath.Glob("templates/widgets/*.html")
	layouts = append(layouts, widgets...)
	layouts = append(layouts, "templates/"+name)

	// parse template every time
	tpl.templates = template.Must(template.ParseFiles(layouts...))
	return tpl.templates.ExecuteTemplate(w, "base", data)
}

func main() {

	// Init
	e := echo.New()

	// Template
	e.Renderer = &Template{}
	//e.Renderer = &Template{templates: template.Must(template.ParseGlob("templates/*/*.html"))}

	// Route
	config.SetRoute(e)

	// Start
	port := os.Getenv("APP_PORT")
	e.Logger.Fatal(e.Start(":" + port))
}
