package api

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gookit/validate"
	"github.com/hmada15/go-food-order/database"
	"github.com/hmada15/go-food-order/helper"
	"github.com/hmada15/go-food-order/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var db *sql.DB = database.Conn()
var ctx = context.Background()

func ListUser(c echo.Context) error {
	users, _ := models.Users().All(ctx, db)

	response := map[string]models.UserSlice{
		"data": users,
	}

	return c.JSON(http.StatusOK, response)
}

func GetUser(c echo.Context) error {
	id := helper.StrToInt(c.Param("id"))
	user, _ := models.FindUser(ctx, db, id)
	if user == nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	response := map[string]*models.User{
		"data": user,
	}

	return c.JSON(http.StatusOK, response)
}

func CreateUser(c echo.Context) error {
	validate.AddValidator("emailUnique", func(val interface{}) bool {
		user, _ := models.Users(qm.Where("email=?", c.FormValue("email"))).One(ctx, db)
		return user == nil
	})

	data, err := validate.FromRequest(c.Request())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	v := data.Create()
	v.StringRule("name", "required|string|minLen:6")
	v.StringRule("email", "required|string|email")
	v.AddRule("email", "emailUnique").SetMessage("Email Must be unique")
	v.StringRule("password", "required|minLen:7|maxLen:20")

	if v.Validate() {
		var user models.User
		user.Name = c.FormValue("name")
		user.Email = c.FormValue("email")
		user.Password = c.FormValue("password")
		err := user.Insert(ctx, db, boil.Infer())
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		response := map[string]interface{}{
			"message": "User created successfully",
			"data":    user,
		}

		return c.JSON(http.StatusOK, response)
	} else {
		return c.JSON(http.StatusBadRequest, v.Errors)
	}
}

func UpdateUser(c echo.Context) error {
	id := helper.StrToInt(c.Param("id"))

	validate.AddValidator("emailUnique", func(val interface{}) bool {
		user, _ := models.Users(models.UserWhere.Email.EQ(c.FormValue("email")), models.UserWhere.ID.NEQ(id)).One(ctx, db)
		return user == nil
	})

	data, err := validate.FromRequest(c.Request())
	if err != nil {
		panic(err)
	}

	v := data.Create()
	v.StringRule("name", "required|string|minLen:6")
	v.StringRule("email", "required|string|email")
	v.AddRule("email", "emailUnique").SetMessage("Email Must be unique")
	v.StringRule("password", "minLen:7|maxLen:20")

	if v.Validate() {
		user, _ := models.FindUser(ctx, db, id)
		user.Name = c.FormValue("name")
		user.Email = c.FormValue("email")
		if c.FormValue("password") != "" {
			user.Password = c.FormValue("password")
		}

		_, err := user.Update(ctx, db, boil.Infer())
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		response := map[string]interface{}{
			"message": "User updated successfully",
			"data":    user,
		}

		return c.JSON(http.StatusOK, response)
	} else {
		return c.JSON(http.StatusBadRequest, v.Errors)
	}
}

func DeleteUser(c echo.Context) error {
	id := helper.StrToInt(c.Param("id"))
	user, _ := models.FindUser(ctx, db, id)
	if user == nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	_, err := user.Delete(ctx, db)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	response := map[string]string{
		"message": "User deleted successfully",
	}

	return c.JSON(http.StatusOK, response)
}
