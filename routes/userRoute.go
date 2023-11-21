package routes

import (
	"github.com/kaleabbyh/golang-santim-echo/controllers"
	"github.com/labstack/echo/v4"
)

func SetupUserRoutes(router *echo.Echo) {
	router.POST("/user/register", controllers.RegisterUser)
	router.POST("/user/login", controllers.LoginUser)
	router.GET("/users", controllers.GetUsers)
}