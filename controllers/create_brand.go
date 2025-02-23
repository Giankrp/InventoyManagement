package controllers

import (
	"github.com/Giankrp/inventoryManagement/db"
	"github.com/Giankrp/inventoryManagement/models"
)

func CreateBrand(name string) (*models.Brand, error) {
	brand := models.Brand{Name: name}

	if err := db.Db.Create(&brand).Error; err != nil {
		return nil, err
	}
	return &brand, nil
}
