package main

import (
	"alterra-agmc/day-3/config"
	"alterra-agmc/day-3/middlewares"
	"alterra-agmc/day-3/routes"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)



func main() {

	var db  *gorm.DB   = config.OpenConnection()
	defer config.CloseConnection(db)
	
	e := echo.New()
	e.Debug = true

	middlewares.Log(e)
	routes.InitRoute(e, db)
	
	e.Logger.Fatal(e.Start(":8080"))
}