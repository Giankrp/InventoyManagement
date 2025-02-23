package controllers

import (
	"math"

	"github.com/Giankrp/inventoryManagement/db"
	"github.com/Giankrp/inventoryManagement/models"
)

func GetProductsPaginated(page, limit int) (models.PaginatedResponse, error) {
	var products []models.Product

	var totalItems int64

	if err := db.Db.Model(&models.Product{}).Count(&totalItems).Error; err != nil {
		return models.PaginatedResponse{}, err
	}

	offSet := (page - 1) * limit

	if err := db.Db.Preload("Brand").Limit(limit).Offset(offSet).Find(&products).Error; err != nil {
		return models.PaginatedResponse{}, err
	}

	var productResponse []models.ProductResponse

	for _, v := range products {
		productResponse = append(productResponse, models.ProductResponse{
			Name:        v.Name,
			Price:       v.Price,
			Description: v.Description,
			Stock:       v.Stock,
			Barcode:     v.Barcode,
		})
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	return models.PaginatedResponse{
		Data:       productResponse,
		Page:       page,
		Limit:      limit,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}, nil
}
