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
			donor.GET("/get/:id", controller.GetDonorById)
			donor.GET("/login/:email/:password", controller.LoginDonor)
			donor.POST("/create", controller.CreateDonor)
			donor.PUT("/update/:id", controller.UpdateDonor)
			donor.DELETE("/delete/:id", controller.DeleteDonor)
		}
		post := main.Group("/post")
		{
			post.GET("/list", controller.ListPosts)
			post.GET("/get/:id", controller.GetPostById)
			post.POST("/create", controller.CreatePost)
			post.PUT("/update/:id", controller.UpdatePostDescription)
			post.DELETE("/delete/:id", controller.DeleteDonor)
		}
	}
	e.Logger.Fatal(e.Start(":8080"))
}
