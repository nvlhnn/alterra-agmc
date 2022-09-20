package factory

import (
	"alterra-agmc/day-6/internal/repository"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Factory struct {
	UserRepository repository.User
	BookRepository repository.Book
}

func NewFactory(mysqlDb *gorm.DB, mongoDB *mongo.Database) *Factory {

	bookCollection := mongoDB.Collection("book")
	return &Factory{
		repository.NewUser(mysqlDb),
		repository.NewBook(bookCollection),
	}
}