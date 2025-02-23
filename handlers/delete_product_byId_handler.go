package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Giankrp/inventoryManagement/controllers"
)

func DeleteProductById(c echo.Context) error {
	productId := c.Param("id")

	err := controllers.DeleteProductById(productId)
	if err != nil {
		if err.Error() == "No se pudo encontrar el producto" {
			c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Producto eliminado"})
}
