package controllers

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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
		return echo.NewHTTPError(http.StatusInternalServerError,err)
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


func GetMyAccounts(c echo.Context) error {
	userID := c.Get("userID")

	if userID=="" {
		return echo.NewHTTPError(http.StatusBadRequest, "not authenticated")
	}

	var accounts []Account
	result := db.Find(&accounts,"user_id=?",userID)
    if result.Error != nil {
        return echo.NewHTTPError(http.StatusBadRequest,  result.Error)       
    }

	var user User
	result = db.First(&user,userID)
    if result.Error != nil {
        return echo.NewHTTPError(http.StatusBadRequest,  result.Error)       
    }

	type Response struct {
		Accounts []Account `json:"accounts"`
		User     User      `json:"user"`
	}

	response := Response{
		Accounts: accounts,
		User:     user,
	}


	return c.JSON(http.StatusCreated, response)
}


func GetAllAcounts(c echo.Context) error {
	userID := c.Get("userID")

	if userID=="" {
		return echo.NewHTTPError(http.StatusBadRequest, "not authenticated")
	}

	var accounts []Account
	result := db.Find(&accounts)
    if result.Error != nil {
        return echo.NewHTTPError(http.StatusBadRequest,  result.Error)       
    }

	return c.JSON(http.StatusCreated, accounts)
}

func DeleteAccount(c echo.Context) error {
	userID := c.Get("userID")

	if userID == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "not authenticated")
	}

	accountID := c.Param("id")
	var account Account
	result := db.First(&account, "id = ?", accountID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "account not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}

	result = db.Delete(&account)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.NoContent(http.StatusNoContent)
}

func UpdateAccount(c echo.Context) error {
	userID := c.Get("userID")

	if userID == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "not authenticated")
	}

	accountID := c.Param("id")

	var account Account
	result := db.First(&account, "id = ? AND user_id = ?", accountID, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "account not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	
	if err := c.Bind(&account); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result = db.Save(&account)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}

	return c.JSON(http.StatusOK, account)
}