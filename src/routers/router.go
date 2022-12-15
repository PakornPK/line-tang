package routers

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/pakornpk/line-tang/src/controllers"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Router(e *echo.Echo) {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/html/*.html")),
	}
	e.Renderer = t

	e.GET("/", controllers.HomePages)
}
