package controllers

import (
	"alterra-agmc/day-4/lib/database"
	"alterra-agmc/day-4/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	bookE *echo.Echo
	bookRepo database.BookDatabase
	bookController BookController
	
)

func seedBookDB()  {
	bookRepo.Create( models.Book{
			Name: "book1",
			Author: "book author",
		})



}

func cleanBookDb()  {
	models.Books = []models.Book{}
}

// func TestMain(m *testing.M) {
//     os.Setenv("ENV", "testing")
// 	bookE = echo.New()
// 	bookRepo = database.NewBook()
// 	bookController = NewBookController(bookRepo)
//     code := m.Run() 
//     os.Exit(code)
// }

func TestCreateBook_success(t *testing.T)  {
	seedBookDB()

	user := models.Book{
		Name:    "booktest",
		Author: "author",
	}

	data, _ := json.Marshal(user)
	reader := bytes.NewReader(data)
	
	req := httptest.NewRequest(http.MethodPost, "/", reader)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books")


	if assert.NoError(t, bookController.CreateBooks(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)	
	}

	cleanBookDb()
}

func TestGetBooks_success(t *testing.T)  {
	seedBookDB()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()

	c := bookE.NewContext(req, rec)
	c.SetPath("/v1/books")


	if assert.NoError(t, bookController.GetBooks(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)	
	}

	cleanBookDb();
}

func TestGetBookById_success(t *testing.T)  {
	seedBookDB()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := bookE.NewContext(req, rec)
	c.SetPath("/v1/books/:id")

	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, bookController.GetBookById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)	
	}

	cleanBookDb()
}

func TestGetBookById_fail(t *testing.T)  {
	seedBookDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books/:id")

	c.SetParamNames("id")
	c.SetParamValues("10")


	if assert.NoError(t, bookController.GetBookById(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)	
	}

	cleanBookDb()
}




func TestUpdateBook_success(t *testing.T)  {
	seedBookDB()

	user := models.Book{
		Name:    "booktest",
		Author: "author",
	}

	data, _ := json.Marshal(user)
	reader := bytes.NewReader(data)
	
	req := httptest.NewRequest(http.MethodPut, "/", reader)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books/:id")

	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, bookController.UpdateBook(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)	
	}

	cleanBookDb()
}

func TestUpdateBook_fail(t *testing.T)  {
	seedBookDB()
	user := models.Book{
		Name:    "booktest",
		Author: "author",
	}

	data, _ := json.Marshal(user)
	reader := bytes.NewReader(data)
	
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", reader)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books/:id")

	c.SetParamNames("id")
	c.SetParamValues("10")


	if assert.NoError(t, bookController.UpdateBook(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)	
	}

	cleanBookDb()
}

func TestDeleteBook_success(t *testing.T)  {
	seedBookDB()

	
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books/:id")

	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, bookController.DeleteBook(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)	
	}

	cleanBookDb()
}

func TestDeleteBook_fail(t *testing.T)  {
	seedBookDB()

	
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books/:id")

	c.SetParamNames("id")
	c.SetParamValues("10")

	if assert.NoError(t, bookController.DeleteBook(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)	
	}

	cleanBookDb()
}