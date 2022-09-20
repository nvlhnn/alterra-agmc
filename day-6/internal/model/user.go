package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}


func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.HashPassword()
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.HashPassword()
	return
}

func (u *User) HashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(bytes)
}

