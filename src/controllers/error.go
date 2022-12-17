package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorPages(c echo.Context, err map[string]interface{}) error {
	return c.Render(http.StatusNotFound, "error.html", err)
}
