package auth

import (
	"alterra-agmc/day-6/internal/factory"
	"alterra-agmc/day-6/internal/model"
	"alterra-agmc/day-6/pkg/utils"
	"fmt"
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service *service
}

var err error

func NewHandler(f *factory.Factory) *handler {
	service := NewService(f)
	return &handler{service}
}

func (h *handler) Login(c echo.Context) error {

	var user model.User
	
	if err = c.Bind(&user); err != nil {
		return utils.FailedResponse(c, http.StatusBadRequest, "Invalid input")
	}

	if err = c.Validate(user); err != nil {
		return utils.FailedResponse(c, http.StatusBadRequest, "Invalid input")
	}

	token, data, err := h.service.Login(user)
	fmt.Println(err != nil)

	if err != nil {
		return utils.FailedResponse(c, http.StatusInternalServerError, nil)
	}

	if reflect.ValueOf(data).IsZero() || token == "" {
		return utils.FailedResponse(c, http.StatusNotFound, "Invalid email or password")
	}

	return utils.SuccessResponse(c, http.StatusOK, "Login success", echo.Map{"token":token})
}

func (h *handler) Register(c echo.Context) error {

	var user model.User

	if err = c.Bind(&user); err != nil {
		return utils.FailedResponse(c, http.StatusBadRequest, "Invalid input")
	}
	
	if err = c.Validate(user); err != nil {
		return utils.FailedResponse(c, http.StatusBadRequest, "Invalid input")
	}

	data, err := h.service.Register(user)
	if err != nil {
		return utils.FailedResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.SuccessResponse(c, http.StatusCreated, nil, data)
}
