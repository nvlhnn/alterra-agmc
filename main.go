package main

import (
	"alterra-agmc/day-6/database"
	"alterra-agmc/day-6/internal/factory"
	"alterra-agmc/day-6/internal/http"
	"alterra-agmc/day-6/internal/middlewares"
	"alterra-agmc/day-6/pkg/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// load env configuration
func init() {
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}
}

func main() {

	// create postgre instance
	postgre_DB := database.NewDatabaseConfigPostgre(database.DatabaseConfigSQL{
		Username:     os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASS"),
		Port:         os.Getenv("DB_PORT"),
		DatabaseName: os.Getenv("DB_NAME"),
		Host: os.Getenv("DB_HOST"),
	})
	defer postgre_DB.DB()


	// create mongo instance
	mongo_DB := database.NewDatabaseConfigMongo()

	f := factory.NewFactory(postgre_DB, mongo_DB)
	e := echo.New()
	middlewares.Log(e)
	log.Println("ok")

	http.NewHttp(e, f)

	e.Logger.Fatal(e.Start(":"+utils.GoDotEnvVariable("PORT")))

}
