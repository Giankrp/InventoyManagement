package controllers

import (
	"errors"

	"gorm.io/gorm"

	"github.com/Giankrp/inventoryManagement/db"
	"github.com/Giankrp/inventoryManagement/models"
)

func CreateProductController(product *models.Product) error {
	var (
		brand            models.Brand
		existingProducts models.Product
	)

	if err := db.Db.First(&brand, product.BrandID).Error; err != nil {
		return errors.New("Brand not found")
	}

	err := db.Db.Where("barcode = ?", product.Barcode).First(&existingProducts).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		if product.Price <= 0 {
			return errors.New("Price must be greater than 0 for new products")
		}
		if err := db.Db.Create(product).Error; err != nil {
			return err
		}
	} else {
		existingProducts.Stock += product.Stock
		if err := db.Db.Save(&existingProducts).Error; err != nil {
			return err
		}
		*product = existingProducts
	}

	if err := db.Db.Preload("Brand").First(product, "id = ?", product.ID).Error; err != nil {
		return err
	}
	return nil
}
