package controllers

import (
	"alterra-agmc/day-3/lib/database"
	"alterra-agmc/day-3/models"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/labstack/echo/v4"
)


type BookController struct{
	database database.BookDatabase
}

func NewBookController(database database.BookDatabase)  BookController {
	return BookController{database: database}
} 


func (c *BookController) GetBooks(e echo.Context) error {

	books := c.database.GetAll()

	if len(books) < 1 {
		return echo.NewHTTPError(http.StatusOK, map[string]interface{}{
			"status": false,
			"message": "No book avaible",
			"data":   nil,
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"message": nil,
		"data":   books,
	})
}

func (c *BookController) GetBookById(e echo.Context) error {

	id := e.Param("id")
	uintId, _ := strconv.Atoi(id)
	book := c.database.GetById(uint(uintId))

	if reflect.ValueOf(book).IsZero() {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"status": false,
			"message": "Book not found",
			"data":   nil,
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"message": nil,
		"data":   book,
	})
}


func(c *BookController)  CreateBooks(e echo.Context) error  {
	var book models.Book
	err := e.Bind(&book)

	err = e.Validate(book)

	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
	}
	book = c.database.Create(book)

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"message": nil,
		"data":   book,
	})
	
}

func (c *BookController) UpdateBook(e echo.Context) error {

	id := e.Param("id")
	uintId, _ := strconv.Atoi(id)

	var book models.Book
	e.Bind(&book)

	book = c.database.Update(uint(uintId), book)

	if reflect.ValueOf(book).IsZero() {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"status": false,
			"message": "Book not found",
			"data":   nil,
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"message": nil,
		"data":   book,
	})
}

func (c *BookController) DeleteBook(e echo.Context) error {

	id := e.Param("id")
	uintId, _ := strconv.Atoi(id)



	if !c.database.Delete(uint(uintId)) {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"status": false,
			"message": "Book not found",
			"data":   nil,
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"message": "book deleted",
		"data":   nil,
	})
}