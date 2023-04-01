package routes

import (
	c "day-13-orm/controllers"
	"day-13-orm/middlewares"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
	e := echo.New()
	middlewares.LoggerMiddleware(e)
	e.POST("/users", c.CreateUserController)
	e.POST("/users/login", c.UserLogin)
	auth := e.Group("")
	auth.Use(middleware.JWT([]byte(os.Getenv("JWT_KEY"))))
	auth.GET("/users", c.GetUsersController)
	auth.GET("/users/:id", c.GetUserController)
	auth.DELETE("/users/:id", c.DeleteUserController)
	auth.PUT("/users/:id", c.UpdateUserController)

	e.GET("/books", c.GetBooksController)
	e.GET("/books/:id", c.GetBookController)
	auth.POST("/books", c.CreateBookController)
	auth.DELETE("/books/:id", c.DeleteBookController)
	auth.PUT("/books/:id", c.UpdateBookController)

	return e
}
