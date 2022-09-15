package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoute(e *echo.Echo, db *gorm.DB) {
	v1 := e.Group("/v1")
	RegisterBookRoute(v1)
	RegisterUserRoute(v1, db)

}