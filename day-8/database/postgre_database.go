package database

import (
	"alterra-agmc/day-6/internal/model"
	"alterra-agmc/day-6/pkg/utils"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfigSQL struct {
	Username     string
	Password     string
	Port         string
	DatabaseName string
	Host 		 string
}

func CreateConnectionPostgre(dbConfig DatabaseConfigSQL) (*gorm.DB, error) {

	ssl := "require"
	
	if utils.GoDotEnvVariable("APP_ENV") != "development" {
		ssl = "disable"
		
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s TimeZone=Asia/Shanghai",dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.Port, dbConfig.DatabaseName, ssl)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return DB, err
}

func NewDatabaseConfigPostgre(dbConfig DatabaseConfigSQL) *gorm.DB {
	DB, err := CreateConnectionPostgre(dbConfig)

	if err != nil {
		log.Fatal(err)
	}

	DB.Debug().AutoMigrate(
		&model.User{},
	)

	return DB
}
