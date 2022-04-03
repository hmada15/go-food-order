package main

import (
	"log"
	"os"

	"github.com/hmada15/go-food-order/helper"
	"github.com/hmada15/go-food-order/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	//get env file values
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	//set debug mode to .env value
	e.Debug = helper.StrToBool(os.Getenv("DEBUG"))

	routes.Routes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
