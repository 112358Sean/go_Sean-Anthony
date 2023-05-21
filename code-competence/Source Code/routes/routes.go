package routes

import (
	"code_competence/configs"
	c "code_competence/controllers"
	m "code_competence/middlewares"
	r "code_competence/repositories"
	s "code_competence/services"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	DB = configs.InitDB()

	JWT = m.NewJWTS()

	userR = r.NewUserRepository(DB)
	userS = s.NewUserService(userR)
	userC = c.NewUserController(userS, JWT)

	itemR = r.NewItemRepository(DB)
	itemS = s.NewItemService(itemR)
	itemC = c.NewItemController(itemS)

	kategoriR = r.NewKategoriRepository(DB)
	kategoriS = s.NewKategoriService(kategoriR)
	kategoriC = c.NewKategoriController(kategoriS)
)

func New() *echo.Echo {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	e := echo.New()

	m.LoggerMiddleware(e)

	auth := e.Group("")
	auth.Use(middleware.JWT([]byte(os.Getenv("JWT_KEY"))))
	auth.GET("/users", userC.GetUsersController)
	auth.GET("/users/:id", userC.GetUserController)
	e.POST("/users", userC.CreateController)
	auth.DELETE("/users/:id", userC.DeleteController)
	auth.PUT("/users/:id", userC.UpdateController)

	auth.GET("/items", itemC.GetItemsController)
	auth.GET("/items/:id", itemC.GetItemController)
	auth.POST("/items", itemC.CreateController)
	auth.DELETE("/items/:id", itemC.DeleteController)
	auth.PUT("/items/:id", itemC.UpdateController)
	auth.GET("/items/category/:id_kategori", itemC.GetItemByCategoryController)
	auth.GET("/items", itemC.GetItemByNameController)

	auth.GET("/kategori", kategoriC.GetKategorisController)
	auth.GET("/kategori/:id", kategoriC.GetKategoriController)
	auth.POST("/kategori", kategoriC.CreateController)
	auth.DELETE("/kategori/:id", kategoriC.DeleteController)
	auth.PUT("/kategori/:id", kategoriC.UpdateController)

	e.POST("/login", userC.LoginController)

	return e
}
