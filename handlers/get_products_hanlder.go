package handlers

import (
	"net/http"

	"github.com/Giankrp/inventoryManagement/controllers"
	"github.com/labstack/echo/v4"
)


func GetProductsHandler(c echo.Context) error{
	products, err := controllers.GetProductsController()
	

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error":"Server error"})
	}



	return c.JSON(http.StatusOK, products)
}
