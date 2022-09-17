package routes

import (
	"alterra-agmc/day-4/controllers"
	"alterra-agmc/day-4/lib/database"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)



func RegisterBookRoute(e *echo.Group){

	database := database.NewBook()
	controller := controllers.NewBookController(database)

	book := e.Group("/books")
	
	book.GET("/:id", controller.GetBookById)
	book.GET("", controller.GetBooks)
	
	book.POST("", controller.CreateBooks, middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	book.PUT("/:id", controller.UpdateBook, middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	book.DELETE("/:id", controller.DeleteBook, middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
}