package handlers

import (
	"net/http"

	"github.com/Giankrp/inventoryManagement/controllers"
	"github.com/labstack/echo/v4"
)

func GetCategoryHandler(c echo.Context) error {
	category, err := controllers.GetCategoryController()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Category not found"})
	}

	return c.JSON(http.StatusOK, category)
}
