package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Giankrp/inventoryManagement/controllers"
	"github.com/Giankrp/inventoryManagement/models"
)

func CreateCategoryHandler(c echo.Context) error {
	var category models.Category

	if err := c.Bind(&category); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{"error": "Error al crear el prodcuto"},
		)
	}

	categoryCreate, err := controllers.CreateCategoryController(category.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, categoryCreate)
}
