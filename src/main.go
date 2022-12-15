package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pakornpk/line-tang/src/configs"
	"github.com/pakornpk/line-tang/src/routers"
)

func main() {
	err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routers.Router(e)
	e.Logger.Fatal(e.Start(configs.GetPort()))
}
