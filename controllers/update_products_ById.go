package controllers

import (
	"errors"

	"github.com/Giankrp/inventoryManagement/db"
	"github.com/Giankrp/inventoryManagement/models"
	"gorm.io/gorm"
)

func UpdateProductById(productId string, updatedProduct *models.Product) (*models.Product, error) {
	var product models.Product

	// Buscar el producto existente por ID
	if err := db.Db.First(&product, "id::text LIKE ?", productId+"%").Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("producto no encontrado")
		}
		return nil, err
	}

	// Actualizar los campos del producto existente
	product.Name = updatedProduct.Name
	product.Price = updatedProduct.Price
	product.Stock = updatedProduct.Stock
	product.BrandID = updatedProduct.BrandID

	// Guardar los cambios en la base de datos
	if err := db.Db.Save(&product).Error; err != nil {
		return nil, errors.New("no se pudo actualizar el producto")
	}

	// Recargar el producto con la marca asociada
	if err := db.Db.Preload("Brand").First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}
