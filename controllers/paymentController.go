package controllers

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreatePayment(c echo.Context) error {

	var payment Payment
	if err := c.Bind(&payment); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	UserID, _ := c.Get("userID").(uuid.UUID)
	payment.UserID=UserID

	var ReceiverAccount Account
	result := db.First(&ReceiverAccount, "account_number = ? ", payment.ReceiverAccount)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "reciever account not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}

	var PayerAccount Account
	result = db.First(&PayerAccount, "account_number = ? ", payment.PayerAccount)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "payer account not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}

	if PayerAccount.Balance-payment.Amount < 25 {
		return echo.NewHTTPError(http.StatusInternalServerError, "no enough balance is available")
	}

	if payment.ReceiverAccount == payment.PayerAccount {
		return echo.NewHTTPError(http.StatusInternalServerError, "self account payment is not possible")
	}

	if err := db.Create(&payment).Error; err != nil {
		return c.JSON(http.StatusInternalServerError,"unable to create payment")
	}

	PayerAccount.Balance -= payment.Amount
	if err := db.Save(&PayerAccount).Error; err!=nil{
		return c.JSON(http.StatusInternalServerError,"unable to update payer account")
	}

	ReceiverAccount.Balance += payment.Amount
	if err := db.Save(&ReceiverAccount).Error; err!=nil{
		return c.JSON(http.StatusInternalServerError,"unable to update reciever account")
	}
	
	payerTransaction := Transaction{
		PaymentID:   payment.ID,
		UserID:      payment.UserID,
		Type:        "payed",
		Amount:      payment.Amount,
		TranferedTo: payment.ReceiverAccount,
	}

	if err:=db.Create(&payerTransaction).Error; err != nil {
		return c.JSON(http.StatusInternalServerError,"unable to update reciever account")
	}
	
	recieverTransaction := Transaction{
		PaymentID:     payment.ID,
		UserID:        ReceiverAccount.UserID,
		Type:          "recieved",
		TranferedFrom: payment.PayerAccount,
		Amount:        payment.Amount,
	}

	if err := db.Create(&recieverTransaction).Error; err != nil {
		return c.JSON(http.StatusInternalServerError,err)
	}

	
	return c.JSON(http.StatusCreated, payment)

}
