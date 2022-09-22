package repository

import (
	"alterra-agmc/day-6/internal/model"

	"gorm.io/gorm"
)

type User interface {
	FindAll() ([]model.User, error)
	FindByID(id uint) (model.User, error)
	Update( user model.User) (model.User, error)
	Delete(id uint) (error)
	Create(user model.User) (model.User, error)
	FindByEmail(email string) (model.User, error)
}

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *user {
	return &user{
		db,
	}
}

func (r *user) FindAll() ([]model.User, error) {
	var users []model.User
	
	err := r.db.Find(&users).Error
	if  err != nil {
		return users, err
	}

	return users, nil
}

func (r *user) FindByID( id uint) (model.User, error) {
	var user model.User

	err := r.db.Find(&user, id).Error
	if  err != nil {
		return user, err
	}

	return user, nil
}

func (r *user) FindByEmail(email string) (model.User, error) {
	var user model.User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if  err != nil {
		return user, err
	}

	return user, nil
}

func (r *user) Update(user model.User) (model.User, error) {
	
	err := r.db.Updates(&user).Error
	if  err != nil {
		return user, err
	}

	return user, err
}


func (r *user) Delete(id uint) error  {

	var user model.User
	err := r.db.Where("id = ?", id).Unscoped().Delete(&user).Error
	if  err != nil {
		return err
	}

	return nil
}

func (r *user) Create(user model.User) (model.User, error)  {
	
	err := r.db.Save(&user).Error
	if  err != nil {
		return user, err
	}

	return user, err
}
