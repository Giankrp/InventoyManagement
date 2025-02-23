package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Giankrp/inventoryManagement/controllers"
)

func GetProductByIdHandler(c echo.Context) error {
	productParam := c.Param("id")

	if len(productParam) < 4 {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{"error": "El id debe tener como minimo 4 caracteres"},
		)
	}
	product, err := controllers.GetProductsById(productParam)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, product)
}
