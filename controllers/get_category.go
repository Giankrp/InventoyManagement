package controllers

import (
	"github.com/Giankrp/inventoryManagement/db"
	"github.com/Giankrp/inventoryManagement/models"
)

func GetCategoryController() ([]models.Category, error ) {
	var categories []models.Category
	if err := db.Db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}
