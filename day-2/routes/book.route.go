package routes

import (
	"alterra-agmc/day-2/controllers"
	"alterra-agmc/day-2/lib/database"

	"github.com/labstack/echo/v4"
)



func RegisterBookRoute(e *echo.Group){

	database := database.NewBook()
	controller := controllers.NewBookController(database)

	book := e.Group("/books")

	book.GET("", controller.GetBooks)
	book.GET("/:id", controller.GetBookById)
	book.POST("", controller.CreateBooks)
	book.PUT("/:id", controller.UpdateBook)
	book.DELETE("/:id", controller.DeleteBook)
}