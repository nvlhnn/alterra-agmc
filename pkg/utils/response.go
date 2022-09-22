package utils

import "github.com/labstack/echo/v4"

type response struct {
	Status  bool `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{}`json:"data"`
}

func SuccessResponse(e echo.Context, code int, message interface{}, data interface{}) error {
	jsonRes := response{
		Status:  true,
		Message: message,
		Data:    data,
	}

	return e.JSON(code, jsonRes)
}

func FailedResponse(e echo.Context, code int, message interface{}) error {
	jsonRes := response{
		Status:  false,
		Message: message,
		Data:    nil,
	}

	return e.JSON(code, jsonRes)
}