package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Giankrp/inventoryManagement/controllers"
	"github.com/Giankrp/inventoryManagement/models"
)

func CreateBrandHandler(c echo.Context) error {
	var brand models.Brand

	if err := c.Bind(&brand); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	createBrand, err := controllers.CreateBrand(brand.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	return c.JSON(http.StatusCreated, createBrand)
}
