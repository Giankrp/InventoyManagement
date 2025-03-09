package controllers

import (
	"github.com/Giankrp/inventoryManagement/db"
	"github.com/Giankrp/inventoryManagement/models"
)

func GetProductsController() ([]models.ProductResponse,error) {
	var products []models.Product

	
	if err := db.Db.Preload("Brand").Preload("category").Find(&products).Error; err != nil {
		return nil, err
	}

	var productResponse []models.ProductResponse

	for _, p := range products {
		productResponse = append(productResponse, models.ProductResponse{
			Name: p.Name,
			Price: p.Price,
			Description: p.Description,
			Stock: p.Stock,
			Barcode: p.Barcode,
			BrandName: p.Brand.Name,
			CategoryName: p.Category.Name,
		})
	}
	return productResponse, nil
}
