package routers

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pakornpk/line-tang/src/controllers"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	var errMsg interface{}
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		errMsg = he.Message
	}
	c.Logger().Error(err)
	errorPage := fmt.Sprintf("%d.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
	controllers.ErrorPages(c, map[string]interface{}{
		"errCode": code,
		"errMsg":  errMsg,
	})
}

func Router(e *echo.Echo) {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/html/*.html")),
	}
	e.Renderer = t
	e.Static("/css", "public/views/css")
	e.Static("/js", "public/views/js")
	e.Static("/assets", "public/assets")
	e.HTTPErrorHandler = customHTTPErrorHandler

	e.GET("/", controllers.HomePages)
	e.GET("/search", controllers.SearchPages)
}
