package main

import (
	"github.com/kaleabbyh/golang-santim-echo/routes"
	"github.com/labstack/echo/v4"
)



func main() {

	router := echo.New()
	
	routes.SetupUserRoutes(router)
	routes.SetupAccountRoutes(router)
	
	router.Start(":8080")
	
}