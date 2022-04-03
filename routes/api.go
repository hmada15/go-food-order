package routes

import (
	"github.com/hmada15/go-food-order/api"
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo) {

	prefix := echo.Group("/api")

	prefix.GET("/user", api.ListUser)
	prefix.GET("/user/:id", api.GetUser)
	prefix.POST("/user", api.CreateUser)
	prefix.PUT("/user/:id", api.UpdateUser)
	prefix.DELETE("/user/:id", api.DeleteUser)
}
