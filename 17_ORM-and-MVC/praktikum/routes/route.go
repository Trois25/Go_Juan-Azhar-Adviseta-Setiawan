package routes

import (
	"praktikum/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// Users
	e.GET("/users", controllers.GetUsersController)
	e.POST("/users", controllers.CreateUserController)
	e.GET("/users/:id", controllers.GetUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)

	// Books
	e.GET("/books", controllers.GetBooksController)
	e.POST("/books", controllers.CreateBookController)
	e.GET("/books/:id", controllers.GetBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)

	// Blogs
	e.GET("/blogs", controllers.GetBlogsController)
	e.POST("/blogs", controllers.CreateBlogController)
	e.GET("/blogs/:id", controllers.GetBlogController)
	e.DELETE("/blogs/:id", controllers.DeleteBlogController)
	e.PUT("/blogs/:id", controllers.UpdateBlogController)

	return e
}
