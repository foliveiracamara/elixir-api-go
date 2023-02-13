package main

import (
	"github.com/foliveiracamara/elixir-api-go/adapter/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	main := e.Group("api/v1")
	{
		donor := main.Group("/donor")
		{
			donor.GET("/list", controller.ListDonors)
			donor.POST("/create", controller.CreateDonor)
		}
	}
	e.Logger.Fatal(e.Start(":8080"))
}
