package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/Giankrp/inventoryManagement/handlers"
)

func Routes(e *echo.Echo) {
	e.POST("/products", handlers.CreateProductsHandler)
	e.POST("/brands", handlers.CreateBrandHandler)
	e.GET("/products", handlers.GetProductsHandler)
	e.GET("products/:id",handlers.GetProductByIdHandler)
	e.GET("/products/paginated", handlers.GetPaginatedProductHandler)
	e.GET("/brands/:id", handlers.GetBrandIdHandler)
	e.PUT("/products/:id", handlers.UpdateProductHandler)
	e.DELETE("/products/:id", handlers.DeleteProductById)
}
