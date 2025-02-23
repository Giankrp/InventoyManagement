package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/Giankrp/inventoryManagement/controllers"
)

func GetBrandIdHandler(c echo.Context) error {
	brandId := c.Param("id")

	brandID, err := strconv.Atoi(brandId)
	if err != nil {
		return err
	}

	id, err := controllers.GetBrandId(uint(brandID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Brand not found"})
	}

	return c.JSON(http.StatusOK, id)
}
