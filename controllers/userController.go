package controllers

import (
	"net/http"
	"strings"

	"github.com/kaleabbyh/golang-santim-echo/utils"
	"github.com/labstack/echo/v4"
)

func RegisterUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if user.Email == "" || user.Name == "" || user.Password == "" || user.Role == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Email, name, password, and role are required fields")
	}
	user.Password,_ = utils.HashPassword(user.Password)
	if err := db.Create(user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
	}

	token,_ := utils.GenerateToken(user.ID, user.Role)
	newUserResponse :=UserResponse{
		Status 		:http.StatusOK,
		User   		: user,
		Token  		: token,
	}

	return c.JSON(http.StatusOK, newUserResponse)
}




func LoginUser(c echo.Context) error{
	loginData := new(LoginData)
    err := c.Bind(&loginData)
    if err != nil {
		 echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")   
    }

	user := new(User)
    result := db.First(&user, "email = ?", strings.ToLower(loginData.Email))
    if result.Error != nil {
        return echo.NewHTTPError(http.StatusBadRequest,  "Invalid email")
        
    }
    token, _ := utils.GenerateToken(user.ID,user.Role)
    if err := utils.VerifyPassword(user.Password, loginData.Password); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,  "Invalid password")
        
    }

	newUserResponse :=UserResponse{
		Status 		:http.StatusOK,
		User   		: user,
		Token  		: token,
	}

	return c.JSON(http.StatusOK, newUserResponse)
}


func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, "Users")
}

func GetUser(c echo.Context) error {
	return c.JSON(http.StatusNotFound, "User not found")
}