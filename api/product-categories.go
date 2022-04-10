package api

import (
	"net/http"

	"github.com/gookit/validate"
	"github.com/hmada15/go-food-order/database"
	"github.com/hmada15/go-food-order/helper"
	"github.com/hmada15/go-food-order/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func ListCategory(c echo.Context) error {
	categories, _ := models.ProductCategories().All(ctx, database.DB)

	response := map[string]models.ProductCategorySlice{
		"data": categories,
	}

	return c.JSON(http.StatusOK, response)
}

func CreateCategory(c echo.Context) error {
	data, err := validate.FromRequest(c.Request())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	v := data.Create()
	v.StringRule("name", "required|string")
	v.StringRule("description", "required|string")
	v.StringRule("is_publish", "required")

	if v.Validate() {
		var category models.ProductCategory
		category.Name = c.FormValue("name")
		category.Slug = helper.StrToSlug(c.FormValue("name"))
		category.Description = null.String{c.FormValue("description"), true}
		category.IsPublish = helper.StrToBool(c.FormValue("is_publish"))
		err := category.Insert(ctx, database.DB, boil.Infer())
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		response := map[string]interface{}{
			"message": "Product Category created successfully",
			"data":    category,
		}

		return c.JSON(http.StatusOK, response)
	} else {
		return c.JSON(http.StatusBadRequest, v.Errors)
	}
}

func GetCategory(c echo.Context) error {
	id := helper.StrToInt(c.Param("id"))
	category, _ := models.FindProductCategory(ctx, database.DB, id)
	if category == nil {
		return c.JSON(http.StatusNotFound, "Category not found")
	}

	response := map[string]*models.ProductCategory{
		"data": category,
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateCategory(c echo.Context) error {
	data, err := validate.FromRequest(c.Request())
	if err != nil {
		panic(err)
	}

	v := data.Create()
	v.StringRule("name", "required|string")
	v.StringRule("description", "required|string")
	v.StringRule("is_publish", "required")

	if v.Validate() {
		id := helper.StrToInt(c.Param("id"))
		category, _ := models.FindProductCategory(ctx, database.DB, id)
		if category == nil {
			return c.JSON(http.StatusNotFound, "Category not found")
		}

		category.Name = c.FormValue("name")
		category.Slug = helper.StrToSlug(c.FormValue("name"))
		category.Description = null.String{c.FormValue("description"), true}
		category.IsPublish = helper.StrToBool(c.FormValue("is_publish"))

		_, err := category.Update(ctx, database.DB, boil.Infer())
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		response := map[string]interface{}{
			"message": "Category updated successfully",
			"data":    category,
		}

		return c.JSON(http.StatusOK, response)
	} else {
		return c.JSON(http.StatusBadRequest, v.Errors)
	}
}

func DeleteCategory(c echo.Context) error {
	id := helper.StrToInt(c.Param("id"))
	category, _ := models.FindProductCategory(ctx, database.DB, id)
	if category == nil {
		return c.JSON(http.StatusNotFound, "Category not found")
	}

	_, err := category.Delete(ctx, database.DB)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	response := map[string]string{
		"message": "Category deleted successfully",
	}

	return c.JSON(http.StatusOK, response)
}
