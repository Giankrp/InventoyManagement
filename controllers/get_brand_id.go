package controllers

import (
	"errors"

	"gorm.io/gorm"

	"github.com/Giankrp/inventoryManagement/db"
	"github.com/Giankrp/inventoryManagement/models"
)

func GetBrandId(id uint) (*models.Brand, error) {
	var brandId models.Brand

	if err := db.Db.First(&brandId,id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Marca no encontrada")
		}
	}
	return &brandId, nil
}
