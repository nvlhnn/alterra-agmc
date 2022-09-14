package controllers

import (
	"alterra-agmc/day-2/lib/database"
	"alterra-agmc/day-2/models"
	"net/http"
	"reflect"
	"strconv"

	"github.com/labstack/echo/v4"
)


type UserController struct{
	database database.UserDatabase
}

func NewUserController(database database.UserDatabase)  UserController {
	return UserController{database: database}
} 


func (c *UserController) GetUsers(e echo.Context) error {

	users, err := c.database.GetAll()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"status": false,
			"message": "Internal server error",
			"data":   nil,
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"message": nil,
		"data":   users,
	})
}

func (c *UserController) GetUserById(e echo.Context) error {

	id := e.Param("id")
	uintId, _ := strconv.Atoi(id)

	user, err := c.database.GetById(uint(uintId))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"status": false,
			"message": "Internal server error",
			"data":   nil,
		})
	}

	if reflect.ValueOf(user).IsZero() {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"status": false,
			"message": "User not found",
			"data":   nil,
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"message": nil,
		"data":   user,
	})
}


func(c *UserController)  CreateUser(e echo.Context) error  {
	var user models.User
	e.Bind(&user)

	user, err := c.database.Create(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"status": false,
			"message": "Internal server error",
			"data":   nil,
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"message": nil,
		"data":   user,
	})
	
}

func (c *UserController) UpdateUser(e echo.Context) error {

	id := e.Param("id")
	uintId, _ := strconv.Atoi(id)

	var updatedUser models.User
	e.Bind(&updatedUser)

	user, err := c.database.GetById(uint(uintId))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"status": false,
			"message": "Internal server error",
			"data":   nil,
		})
	}

	if reflect.ValueOf(user).IsZero() {
		return echo.NewHTTPError(http.StatusBadGateway, map[string]interface{}{
			"status": false,
			"message": "Invalid id",
			"data":   nil,
		})
	}

	if updatedUser.Email != "" {
		user.Email = updatedUser.Email
	}

	if updatedUser.Password != "" {
		user.Password = updatedUser.Password
	}

	user, err = c.database.Update(uint(uintId), user)
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, map[string]interface{}{
			"status": false,
			"message": err.Error(),
			"data":   nil,
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"message": nil,
		"data":   user,
	})
}

func (c *UserController) DeleteUser(e echo.Context) error {

	id := e.Param("id")
	uintId, _ := strconv.Atoi(id)

	user, err := c.database.GetById(uint(uintId))
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, map[string]interface{}{
			"status": false,
			"message": "Internal server error",
			"data":   nil,
		})
	}

	if reflect.ValueOf(user).IsZero() {
		return echo.NewHTTPError(http.StatusBadGateway, map[string]interface{}{
			"status": false,
			"message": "Invalid id",
			"data":   nil,
		})
	}

	err = c.database.Delete(uint(uintId))
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, map[string]interface{}{
			"status": false,
			"message": err.Error(),
			"data":   nil,
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"message": "User deleted",
		"data":   nil,
	})
}