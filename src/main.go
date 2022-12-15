package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pakornpk/line-tang/src/routers"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	routers.Router(e)
	e.Logger.Fatal(e.Start(":1323"))
}
