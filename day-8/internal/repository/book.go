package repository

import (
	"alterra-agmc/day-6/internal/model"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Book interface {
	FindAll() ([]model.Book, error)
	// FindByID(id uint) (model.User, error)
	// Update( user model.User) (model.User, error)
	// Delete(id uint) (error)
	// Create(user model.User) (model.User, error)
	// FindByEmail(email string) (model.User, error)
}

type book struct {
	cl *mongo.Collection
}

func NewBook(cl *mongo.Collection) *book {
	return &book{
		cl,
	}
}


func (db *book) FindAll() ([]model.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var books []model.Book
	
	results, err :=  db.cl.Find(ctx, bson.M{})
	if  err != nil {
		return books, err
	}

	defer results.Close(ctx)

	for results.Next(ctx) {
		var book model.Book
        err := results.Decode(&book)
        if err != nil {
            log.Fatal(err.Error())
        }

        books = append(books, book)
	}

	return books, nil
}

// func (r *user) FindByID( id uint) (model.User, error) {
// 	var user model.User

// 	err := r.db.Find(&user, id).Error
// 	if  err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }

// func (r *user) FindByEmail(email string) (model.User, error) {
// 	var user model.User

// 	err := r.db.Where("email = ?", email).Find(&user).Error
// 	if  err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }

// func (r *user) Update(user model.User) (model.User, error) {
	
// 	err := r.db.Updates(&user).Error
// 	if  err != nil {
// 		return user, err
// 	}

// 	return user, err
// }


// func (r *user) Delete(id uint) error  {

// 	var user model.User
// 	err := r.db.Where("id = ?", id).Unscoped().Delete(&user).Error
// 	if  err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (r *user) Create(user model.User) (model.User, error)  {
	
// 	err := r.db.Save(&user).Error
// 	if  err != nil {
// 		return user, err
// 	}

// 	return user, err
// }
