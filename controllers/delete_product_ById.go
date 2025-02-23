package controllers

import (
	"errors"

	"github.com/Giankrp/inventoryManagement/db"
	"github.com/Giankrp/inventoryManagement/models"
	"gorm.io/gorm"
)

func DeleteProductById(prductId string) error  {
	
	var product models.Product

	if err := db.Db.First(&product, "id::text LIKE ?", prductId+"%").Error; err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound){
			return errors.New("No se pudo encontrar el producto")
		}
	}

	if err := db.Db.Delete(&product).Error; err != nil{
		return errors.New("No se pudo eliminar el producto")
	}

	return nil
}
