package controllers

import (
	"errors"

	"gorm.io/gorm"

	"github.com/Giankrp/inventoryManagement/db"
	"github.com/Giankrp/inventoryManagement/models"
)

func GetProductsById(productId string) (*models.Product, error) {
	var product models.Product

	if err := db.Db.Preload("Brand").Where("id::text LIKE ?", productId+"%").First(&product).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("producto no encontrado")
		}

		return nil, err
	}
	return &product, nil
}
