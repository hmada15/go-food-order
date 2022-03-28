package main

import (
	"github.com/hmada15/go-food-order/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	
	routes.Routes(e)
	
	e.Logger.Fatal(e.Start(":1323"))
}
