package http

import (
	"alterra-agmc/day-6/internal/app/auth"
	"alterra-agmc/day-6/internal/app/book"
	"alterra-agmc/day-6/internal/app/user"
	"alterra-agmc/day-6/internal/factory"
	"alterra-agmc/day-6/internal/middlewares"
	"encoding/json"
	"io/ioutil"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	e.Validator = &middlewares.CustomValidator{Validator: validator.New()}

	e.GET("/status", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "OK"})
	})

	v1 := e.Group("/v1")

	user.NewHandler(f).Route(v1.Group("/users"))
	auth.NewHandler(f).Route(v1.Group("/auth"))
	book.NewHandler(f).Route(v1.Group("/books"))

	data, err := json.MarshalIndent(e.Routes(), "", "  ")

	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("routes.json", data, 0644)

}