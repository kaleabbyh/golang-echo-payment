package middlewares

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/kaleabbyh/golang-santim-echo/utils"
	"github.com/labstack/echo/v4"
)


func IsLoggedIn(next echo.HandlerFunc) echo.HandlerFunc {

    return func(c echo.Context) error {
	userID, role, err := utils.GetValuesFromToken(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized,  "User not logged in")			
		}

		if userID == uuid.Nil {
			return c.JSON(http.StatusBadRequest,  " Invalid token")
		}

		c.Set("userID", userID)
		c.Set("role", role)
		return next(c)
    }
}