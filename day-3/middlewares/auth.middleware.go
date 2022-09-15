package middlewares

import (
	"alterra-agmc/day-3/utils"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func IsUser(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		paramId := c.Param("id")
		paramIdNumber, _ := strconv.Atoi(paramId)
	
		if user.Valid {
			claims := user.Claims.(jwt.MapClaims)
			userId := claims["userId"].(float64)
			if int(userId) == paramIdNumber {
				return next(c)
			}
		}

		return utils.FailedResponse(c, 403, "Unauthorized")

    }
}