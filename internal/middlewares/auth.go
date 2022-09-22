package middlewares

import (
	"alterra-agmc/day-6/pkg/utils"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId int) (string, error) {	
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"]=time.Now().Add(time.Hour*1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}


func ExtractToken(e echo.Context) uint  {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(uint)
		return userId
	}
	return 0
}

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
