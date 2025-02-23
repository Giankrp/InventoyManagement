package handlers

import (
	"net/http"
	"strconv"

	"github.com/Giankrp/inventoryManagement/controllers"
	"github.com/labstack/echo/v4"
)


func GetPaginatedProductHandler(c echo.Context) error  {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	page, err := strconv.Atoi(pageParam)

	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitParam)

	if err != nil || limit < 1{
		limit = 10
	}


	products, err := controllers.GetProductsPaginated(page,limit)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Sever error"})
	}

	return c.JSON(http.StatusOK, products)
}
