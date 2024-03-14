package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func echoStart() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, "success")
	})

	return e
}

func main() {
	e := echoStart()

	if err := e.Start(":8080"); err != nil {
		panic(err)
	}
}
