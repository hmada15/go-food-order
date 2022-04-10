package routes

import (
	"github.com/hmada15/go-food-order/api"
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo) {

	prefix := echo.Group("/api")

	prefix.GET("/user", api.ListUser)
	prefix.POST("/user", api.CreateUser)
	prefix.GET("/user/:id", api.GetUser)
	prefix.PUT("/user/:id", api.UpdateUser)
	prefix.DELETE("/user/:id", api.DeleteUser)

	prefix.GET("/product-categories", api.ListCategory)
	prefix.POST("/product-categories", api.CreateCategory)
	prefix.GET("/product-categories/:id", api.GetCategory)
	prefix.PUT("/product-categories/:id", api.UpdateCategory)
	prefix.DELETE("/product-categories/:id", api.DeleteCategory)

}
