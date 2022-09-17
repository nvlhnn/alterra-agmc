package database

import (
	"alterra-agmc/day-4/models"
)


type BookDatabase struct{
	books []models.Book
}


func NewBook() BookDatabase  {
	return BookDatabase{models.Books}
}

func (db *BookDatabase) GetAll() []models.Book {
	return db.books
}

func (db *BookDatabase) GetById(id uint) models.Book  {
	
	var book models.Book
	for _, v := range db.books {
		if v.Id == id {
			return v
		}
	}
	return book
}

func (db *BookDatabase) Create(book models.Book) models.Book  {
	
	if len(db.books) > 0 {
		book.Id = db.books[len(db.books) - 1].Id + 1
	}else{
		book.Id = 1
	}

	db.books = append(db.books, book)
	return book
}

func (db *BookDatabase) Update(id uint, updatedBook models.Book) models.Book  {
	
	var book models.Book
	for idx, v := range db.books {
		if v.Id == id {
			if updatedBook.Name != "" {
				db.books[idx].Name = updatedBook.Name
			}

			if updatedBook.Author != "" {
				db.books[idx].Author = updatedBook.Author
			}

			return db.books[idx]
		}
	}
	return book
}


func (db *BookDatabase) Delete(id uint) bool  {
	
	for idx, v := range db.books {
		if v.Id == id {
			db.books = append(db.books[:idx], db.books[idx+1:]... )
			return true
		}
	}
	return false
}



