package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomePages(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", nil)
}
