package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomePages(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", nil)
}

func SearchPages(c echo.Context) error {
	return c.Render(http.StatusOK, "search.html", map[string]interface{}{
		"type": []string{"บางกรอบ", "หนานุ่ม"},
	})
}

func ErrorPages(c echo.Context, err map[string]interface{}) error {
	return c.Render(http.StatusNotFound, "error.html", err)
}
