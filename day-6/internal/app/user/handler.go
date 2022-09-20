package user

import (
	"alterra-agmc/day-6/internal/factory"
	"alterra-agmc/day-6/internal/model"
	"alterra-agmc/day-6/pkg/utils"
	"net/http"
	"reflect"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) Get(c echo.Context) error {

	users, err := h.service.FindAll()

	if err != nil {
		return utils.FailedResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.SuccessResponse(c, 200, nil, users)

}


func (h *handler) GetUser(c echo.Context) error {

	id := c.Param("id")
	uintId, _ := strconv.Atoi(id)

	user, err := h.service.FindByID(uint(uintId))

	if err != nil {
		return utils.FailedResponse(c, http.StatusInternalServerError, nil)
	}

	if reflect.ValueOf(user).IsZero() {
		return utils.FailedResponse(c, http.StatusNotFound, "User not found")
	}

	return utils.SuccessResponse(c, 200, nil, user)

}

func (h *handler) UpdateUser(c echo.Context) error {

	id := c.Param("id")
	uintId, _ := strconv.Atoi(id)

	var updatedUser model.User
	c.Bind(&updatedUser)

	user, err := h.service.Update(uint(uintId), updatedUser)
	if err != nil {
		return utils.FailedResponse(c, http.StatusInternalServerError, nil)
	}

	if reflect.ValueOf(user).IsZero() {
		return utils.FailedResponse(c, http.StatusNotFound, "Invalid id")
	}

	return utils.SuccessResponse(c, 200, nil, user)
}


func (h *handler) DeleteUser(c echo.Context) error {

	id := c.Param("id")
	uintId, _ := strconv.Atoi(id)

	user, err := h.service.Delete(uint(uintId))
	if err != nil {
		return utils.FailedResponse(c, http.StatusInternalServerError, nil)
	}

	if reflect.ValueOf(user).IsZero() {
		return utils.FailedResponse(c, http.StatusNotFound, "Invalid id")
	}

	return utils.SuccessResponse(c, 200, "User deleted", nil)
}