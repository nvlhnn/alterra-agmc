package database

import (
	"alterra-agmc/day-2/models"

	"gorm.io/gorm"
)


type UserDatabase struct{
	db *gorm.DB
}


func NewUser(db *gorm.DB) UserDatabase  {
	return UserDatabase{db}
}

func (qry *UserDatabase) GetAll() ([]models.User, error) {
	var users []models.User
	
	err := qry.db.Find(&users).Error
	if  err != nil {
		return users, err
	}

	return users, nil
}

func (qry *UserDatabase) GetById(id uint) (models.User, error)  {
	
	var user models.User

	err := qry.db.Find(&user, id).Error
	if  err != nil {
		return user, err
	}


	return user, nil
}

func (qry *UserDatabase) Create(user models.User) (models.User, error)  {
	
	err := qry.db.Save(&user).Error
	if  err != nil {
		return user, err
	}

	return user, err
}

func (qry *UserDatabase) Update(id uint, user models.User) (models.User, error) {
	
	err := qry.db.Updates(&user).Error
	if  err != nil {
		return user, err
	}

	return user, err
}


func (qry *UserDatabase) Delete(id uint) error  {

	var user models.User
	err := qry.db.Where("id = ?", id).Delete(&user).Error
	if  err != nil {
		return err
	}

	return nil
}



