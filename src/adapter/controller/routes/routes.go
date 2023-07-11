package routes

import (
	"github.com/foliveiracamara/elixir-api-go/src/adapter/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup, 
	donorController controller.DonorController,
	// postController controller.PostControllerInterface,
) {

	r.POST("api/v1/donor/create", donorController.CreateDonor)
	// r.POST("api/v1/post/create", postController.CreatePost)
	// r.GET("donor/list", controller.ListDonors)
	// r.GET("/get/:id", controller.GetDonorById)
	// r.GET("/login/:email/:password", controller.LoginDonor)
	// r.PUT("/update/:id", controller.UpdateDonor)
	// r.DELETE("/delete/:id", controller.DeleteDonor)
	// r.GET("/list", controller.ListPosts)
	// r.GET("/get/:id", controller.GetPostById)
	// r.POST("/create", controller.CreatePost)
	// r.PUT("/update/:id", controller.UpdatePostDescription)
	// r.DELETE("/delete/:id", controller.DeleteDonor)
}

// e := echo.New()
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())

// 	main := e.Group("api/v1")
// 	{
// 		donor := main.Group("/donor")
// 		{
// 			donor.GET("/list", controller.ListDonors)
// 			donor.GET("/get/:id", controller.GetDonorById)
// 			donor.GET("/login/:email/:password", controller.LoginDonor)
// 			donor.POST("/create", controller.CreateDonor)
// 			donor.PUT("/update/:id", controller.UpdateDonor)
// 			donor.DELETE("/delete/:id", controller.DeleteDonor)
// 		}
// 		post := main.Group("/post")
// 		{
// 			post.GET("/list", controller.ListPosts)
// 			post.GET("/get/:id", controller.GetPostById)
// 			post.POST("/create", controller.CreatePost)
// 			post.PUT("/update/:id", controller.UpdatePostDescription)
// 			post.DELETE("/delete/:id", controller.DeleteDonor)
// 		}
// 	}
// 	e.Logger.Fatal(e.Start(":8080"))