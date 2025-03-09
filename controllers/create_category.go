package controllers

import (
	"github.com/Giankrp/inventoryManagement/db"
	"github.com/Giankrp/inventoryManagement/models"
)

func CreateCategoryController(name string) (*models.Category, error) {
	category := models.Category{Name: name}

	if err:= db.Db.Create(&category).Error; err != nil{
		return nil, err
	}

	return &category, nil
}
