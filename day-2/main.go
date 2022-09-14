package main

import (
	"alterra-agmc/day-2/config"
	"alterra-agmc/day-2/routes"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)




func main() {

	var db  *gorm.DB   = config.OpenConnection()
	defer config.CloseConnection(db)
	
	e := echo.New()
	e.Debug = true

	e.GET("/books", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	routes.InitRoute(e, db)
	
	e.Logger.Fatal(e.Start(":8080"))
}