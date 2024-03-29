package routes

import (
	c "day-13-orm/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// Initialization Echo
	e := echo.New()
	// user routing
	e.GET("/users", c.GetUsersController)
	e.GET("/users/:id", c.GetUserController)
	e.POST("/users", c.CreateUserController)
	e.DELETE("/users/:id", c.DeleteUserController)
	e.PUT("/users/:id", c.UpdateUserController)
	
	e.GET("/books", c.GetBooksController)
	e.GET("/books/:id", c.GetBookController)
	e.POST("/books", c.CreateBookController)
	e.DELETE("/books/:id", c.DeleteBookController)
	e.PUT("/books/:id", c.UpdateBookController)

	return e
}
