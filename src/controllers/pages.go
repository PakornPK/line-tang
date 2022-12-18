package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pakornpk/line-tang/src/services"
)

func HomePages(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", nil)
}

func SearchPages(c echo.Context) error {
	nav, err := services.GetComponent("navbar")
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.html", nil)
	}
	return c.Render(http.StatusOK, "search.html", map[string]interface{}{
		"type": []string{"บางกรอบ", "หนานุ่ม"},
		"nav": nav,
	})
}

func CreatePages(c echo.Context) error {
	nav, err := services.GetComponent("navbar")
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.html", nil)
	}
	return c.Render(http.StatusOK, "create.html", map[string]interface{}{
		"nav": nav,
	})
}

func ErrorPages(c echo.Context, err map[string]interface{}) error {
	return c.Render(http.StatusNotFound, "error.html", err)
}
