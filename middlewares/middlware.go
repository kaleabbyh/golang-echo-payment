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
			return c.JSON(http.StatusUnauthorized,  "Invalid token")			
		}

		if userID == uuid.Nil {
			return c.JSON(http.StatusBadRequest,  "User not logged in")
		}

		c.Set("userID", userID)
		c.Set("role", role)
		return next(c)
    }
}



func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {

    return func(c echo.Context) error {
	userID, role, err := utils.GetValuesFromToken(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized,  "Invalid token")			
		}

		if userID == uuid.Nil {
			return c.JSON(http.StatusBadRequest,  " User not logged in")
		}

		if role != "admin" && role!="superadmin" {
			return c.JSON(http.StatusUnauthorized,  " Unauthorized")
		}

		c.Set("userID", userID)
		c.Set("role", role)
		return next(c)
    }
}