package routes

import (
	"github.com/kaleabbyh/golang-santim-echo/controllers"
	"github.com/kaleabbyh/golang-santim-echo/middlewares"
	"github.com/labstack/echo/v4"
)


func SetupAccountRoutes(router *echo.Echo) {
	router.POST("/account/createaccount", controllers.CreateAccount, middlewares.IsAdmin)
	router.GET("/account/getallaccounts", controllers.GetAllAcounts, middlewares.IsAdmin)
	router.GET("/account/getmyaccounts", controllers.GetMyAccounts, middlewares.IsLoggedIn)
	
	//router.POST("/account/createaccount", controllers.CreateAccount, echo.MiddlewareFunc(middlewares.IsLoggedIn))
}