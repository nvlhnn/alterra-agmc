package user

import (
	"alterra-agmc/day-6/internal/middlewares"
	"alterra-agmc/day-6/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.Get, middleware.JWT([]byte(utils.GoDotEnvVariable("JWT_SECRET"))))
	g.GET("/:id", h.GetUser, middleware.JWT([]byte(utils.GoDotEnvVariable("JWT_SECRET"))))
	g.PUT("/:id", h.UpdateUser, middleware.JWT([]byte(utils.GoDotEnvVariable("JWT_SECRET"))), middlewares.IsUser)
	g.DELETE("/:id", h.DeleteUser, middleware.JWT([]byte(utils.GoDotEnvVariable("JWT_SECRET"))), middlewares.IsUser)
}
