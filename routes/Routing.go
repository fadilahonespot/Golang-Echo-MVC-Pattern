package routes

import (
	"Golang-Echo-MVC-Pattern/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Routing struct {
	example    controller.ExampleController
	user       controller.UserController
	catagory   controller.CatagoryController
	discussion controller.DiscussionController
}

func (Routing Routing) GetRoutes() *echo.Echo {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/posts/", Routing.example.GetPostsController)

	// users
	e.POST("/user/", Routing.user.AddUserController)
	e.GET("/user/", Routing.user.ViewUsersController)
	e.GET("/user/:idUser", Routing.user.ViewUserIdController)

	// catagory
	e.POST("/catagory/", Routing.catagory.AddCatagory)
	e.GET("/catagory/", Routing.catagory.ViewAllCatagory)

	// Discussion
	e.POST("/discussion/", Routing.discussion.AddDiscussion)
	e.POST("/discussion/first/", Routing.discussion.AddDiscussionFirst)
	e.POST("discussion/second/", Routing.discussion.AddDiscussionSecond)
	e.GET("/discussion/", Routing.discussion.GetAllDiscussion)
	e.GET("/discussion/:discussionId", Routing.discussion.GetDiscussionDetailById)
	e.POST("/discussion/upload", Routing.discussion.AddFilesImagesDiscussion)
	e.PUT("/discussion/:discussionId", Routing.discussion.EditDiscussion)
	e.DELETE("/discussion/:discussionId", Routing.discussion.DeleteDiscussionById)

	return e
}
