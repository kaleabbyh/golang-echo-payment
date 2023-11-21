package controllers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateAccount(c echo.Context) error {
	userID := c.Get("userID")

	if userID=="" {
		return echo.NewHTTPError(http.StatusBadRequest, "not authenticated")
	}
	userIDValue, _ := userID.(uuid.UUID)
	
	account := new(Account)
	if err := c.Bind(account); err != nil {
		 return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}
	
	account.CreatedBy =  userIDValue
	if 	account.AccountNumber == "" || 
		account.CreatedBy == uuid.Nil ||
		account.UserID == uuid.Nil || 
		account.Balance == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "AccountNumber, Creater, User, and Balance are required fields")
	}
	
	if err := db.Create(account).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create account", err)
	}

	var user User
	result := db.First(&user, account.UserID)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error.Error())
		
	}

	response := AccountResponse{
		Account: *account,
		User:    user,
		Message: "Account created successfully",
	}

	return c.JSON(http.StatusCreated, response)
}