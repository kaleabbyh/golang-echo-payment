package routes

import (
	"github.com/kaleabbyh/golang-santim-echo/controllers"
	"github.com/kaleabbyh/golang-santim-echo/middlewares"
	"github.com/labstack/echo/v4"
)


func SetupAccountRoutes(router *echo.Echo) {
	router.POST("/account/createaccount", controllers.CreateAccount, echo.MiddlewareFunc(middlewares.IsLoggedIn))
}