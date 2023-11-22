package routes

import (
	"github.com/kaleabbyh/golang-santim-echo/controllers"
	"github.com/kaleabbyh/golang-santim-echo/middlewares"
	"github.com/labstack/echo/v4"
)

func SetupPaymentRoutes(router *echo.Echo) {
	router.POST("/payment/createpayment", controllers.CreatePayment, middlewares.IsLoggedIn)
	//136607294
}