package controllers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateAccount(c echo.Context) error {

	account := new(Account)
	if err := c.Bind(account); err != nil {
		 return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if 	account.AccountNumber == "" || 
		account.CreatedBy == uuid.Nil ||
		account.UserID == uuid.Nil || 
		account.Balance == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "AccountNumber, Creater, User, and Balance are required fields")
	}

	if err := db.Create(account).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create account", err)
	}

	return nil
}