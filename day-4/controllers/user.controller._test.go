package controllers

import (
	"alterra-agmc/day-4/config"
	"alterra-agmc/day-4/lib/database"
	"alterra-agmc/day-4/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	e *echo.Echo
	db *gorm.DB
	repo database.UserDatabase
	controller UserController
	
)

func seedUserDB()  {
	db.AutoMigrate(&models.User{})
	repo.Create(models.User{
		Email: "user1@gmail.com",
		Password: "11111111",
	})
	repo.Create(models.User{
		Email: "user2@gmail.com",
		Password: "11111111",
	})
	repo.Create(models.User{
		Email: "user2@gmail.com",
		Password: "11111111",
	})
}

func cleanUserDb()  {
	// repo.Delete(1)
	// repo.Delete(2)
	// repo.Delete(3)
	db.Migrator().DropTable(&models.User{})
}

func TestMain(m *testing.M) {
    os.Setenv("ENV", "testing")
	e = echo.New()
	db = config.OpenConnection()
	repo = database.NewUser(db)
	controller = NewUserController(repo)

	bookE = echo.New()
	bookRepo = database.NewBook()
	bookController = NewBookController(bookRepo)
	
    code := m.Run() 
    os.Exit(code)
}
func TestGetUsers_success(t *testing.T)  {
	seedUserDB()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/v1/users")


	if assert.NoError(t, controller.GetUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)	
	}

	cleanUserDb();
}

func TestGetUserById_success(t *testing.T)  {
	seedUserDB()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/users/:id")

	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, controller.GetUserById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)	
	}

	cleanUserDb()
}

func TestGetUserById_fail(t *testing.T)  {
	seedUserDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/users/:id")

	c.SetParamNames("id")
	c.SetParamValues("10")


	if assert.NoError(t, controller.GetUserById(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)	
	}

	cleanUserDb()
}

func TestCreateUser_success(t *testing.T)  {
	seedUserDB()

	user := models.User{
		Email:    "usertest@gmail.com",
		Password: "11111111",
	}

	data, _ := json.Marshal(user)
	reader := bytes.NewReader(data)
	
	req := httptest.NewRequest(http.MethodPost, "/", reader)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/users")


	if assert.NoError(t, controller.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)	
	}

	cleanUserDb()
}


func TestUpdateUser_success(t *testing.T)  {
	seedUserDB()

	user := models.User{
		Email:    "usertestupdate@gmail.com",
		Password: "11111111",
	}

	data, _ := json.Marshal(user)
	reader := bytes.NewReader(data)
	
	req := httptest.NewRequest(http.MethodPut, "/", reader)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/users/:id")

	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, controller.UpdateUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)	
	}

	cleanUserDb()
}

func TestUpdateUser_fail(t *testing.T)  {
	seedUserDB()
	user := models.User{
		Email:    "usertestupdate@gmail.com",
		Password: "11111111",
	}

	data, _ := json.Marshal(user)
	reader := bytes.NewReader(data)
	
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", reader)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/users/:id")

	c.SetParamNames("id")
	c.SetParamValues("10")


	if assert.NoError(t, controller.UpdateUser(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)	
	}

	cleanUserDb()
}

func TestDeleteUser_success(t *testing.T)  {
	seedUserDB()

	
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/users/:id")

	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, controller.DeleteUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)	
	}

	cleanUserDb()
}

func TestDeleteUser_fail(t *testing.T)  {
	seedUserDB()

	
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/users/:id")

	c.SetParamNames("id")
	c.SetParamValues("10")

	if assert.NoError(t, controller.DeleteUser(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)	
	}

	cleanUserDb()
}