package models

var Books []Book

type Book struct {
	Id     uint   `json:"id"`
	Name   string `json:"name" binding:"required"`
	Author string `json:"author" binding:"required"`
}
