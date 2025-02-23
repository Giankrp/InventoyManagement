package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/Giankrp/inventoryManagement/controllers"
	"github.com/Giankrp/inventoryManagement/models"
)

func UpdateProductHandler(c echo.Context) error {
	productId := c.Param("id") // Obtener el ID de la URL
	updatedProduct := new(models.Product)

	// Vincular los datos del cuerpo de la solicitud al producto
	if err := c.Bind(updatedProduct); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{"error": "Datos de entrada inválidos"},
		)
	}

	// Validar el producto actualizado
	if err := c.Validate(updatedProduct); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make(map[string]string)
		for _, fieldError := range validationErrors {
			errorMessages[fieldError.Field()] = fieldError.Tag()
		}
		return c.JSON(http.StatusBadRequest, errorMessages)
	}

	// Actualizar el producto en la base de datos
	product, err := controllers.UpdateProductById(productId, updatedProduct)
	if err != nil {
		if err.Error() == "producto no encontrado" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, product)
}
