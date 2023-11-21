package routes

import (
	"github.com/kaleabbyh/golang-santim-echo/controllers"
	"github.com/labstack/echo/v4"
)

func SetupAccountRoutes(router *echo.Echo) {
	router.POST("/user/register", controllers.CreateAcount)
	
}