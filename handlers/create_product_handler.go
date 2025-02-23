package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/Giankrp/inventoryManagement/controllers"
	"github.com/Giankrp/inventoryManagement/models"
)

func CreateProductsHandler(c echo.Context) error {
	product := new(models.Product)

	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// log.Printf("Product after Bind: %+v\n", product)

	if err := c.Validate(product); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make(map[string]string)
		for _, fieldError := range validationErrors {
			errorMessages[fieldError.Field()] = fieldError.Tag()
		}
		return c.JSON(http.StatusBadRequest, errorMessages)
	}
	if err := controllers.CreateProductController(product); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{"error": err.Error()},
		)
	}
	return c.JSON(http.StatusCreated, product)
}
