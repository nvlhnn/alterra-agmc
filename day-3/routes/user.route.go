package routes

import (
	"alterra-agmc/day-3/controllers"
	"alterra-agmc/day-3/lib/database"
	"alterra-agmc/day-3/middlewares"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)



func RegisterUserRoute(e *echo.Group, db *gorm.DB){

	database := database.NewUser(db)
	controller := controllers.NewUserController(database)

	book := e.Group("/users")

	book.POST("", controller.CreateUser)
	book.POST("/login", controller.Login)
	
	// private := e.Group("/users", middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))

	book.GET("", controller.GetUsers, middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	book.GET("/:id", controller.GetUserById, middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	book.PUT("/:id", controller.UpdateUser, middleware.JWT([]byte(os.Getenv("JWT_SECRET"))), middlewares.IsUser)
	book.DELETE("/:id", controller.DeleteUser, middleware.JWT([]byte(os.Getenv("JWT_SECRET"))), middlewares.IsUser)
    // book.Use(middlewareOne)
}
