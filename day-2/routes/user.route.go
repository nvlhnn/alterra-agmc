package routes

import (
	"alterra-agmc/day-2/controllers"
	"alterra-agmc/day-2/lib/database"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)



func RegisterUserRoute(e *echo.Group, db *gorm.DB){

	database := database.NewUser(db)
	controller := controllers.NewUserController(database)

	book := e.Group("/users")

	book.GET("", controller.GetUsers)
	book.GET("/:id", controller.GetUserById)
	book.POST("", controller.CreateUser)
	book.PUT("/:id", controller.UpdateUser)
	book.DELETE("/:id", controller.DeleteUser)
}