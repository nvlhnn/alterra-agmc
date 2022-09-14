package models

var Books []Book

type Book struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}
