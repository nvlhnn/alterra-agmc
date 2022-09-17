package config

import (
	"alterra-agmc/day-4/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	var dbName string
	ssl := "disable"
	env := os.Getenv("ENV")
	_, isTesting := os.LookupEnv("ENV")

	if isTesting {
		err := godotenv.Load("../.env")
		if err != nil {
			panic("failed to load env file")
		}

		dbName = os.Getenv("DB_NAME_TEST")


	}else{
		if env == "production" {
			ssl = "require"
		}else{
			err := godotenv.Load()
			if err != nil {
				panic("failed to load env file")
			}
		}

		dbName = os.Getenv("DB_NAME")
		fmt.Println(dbName)

	}


	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbPort, dbName, ssl)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	fmt.Println("migrating...")
	err = db.AutoMigrate(&models.User{})


	return db
}

func CloseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}