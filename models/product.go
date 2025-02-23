package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	AddedAt   time.Time      `json:"added_at"   gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Product struct {
	BaseModel
	ID          uuid.UUID `json:"id"          gorm:"primaryKey"`
	Name        string    `json:"name"        gorm:"not null"                                                validate:"required"`
	Price       float64   `json:"price"       gorm:"not null"                                                validate:"gte=0"`
	Description string    `json:"description"                                                                validate:"max=500"`
	Stock       int       `json:"stock"       gorm:"not null"                                                validate:"gte=0"`
	Barcode     string    `json:"barcode"     gorm:"not null"                                                validate:"required"`
	BrandID     uint      `json:"brand_id"    gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Brand       Brand     `json:"brand"       gorm:"foreignKey:BrandID;"`
	CategoryID  uint      `json:"category_id" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" validate:"required"` 
	Category    Category  `json:"category"    gorm:"foreignKey:CategoryID;"`
}

type Brand struct {
	ID   uint   `json:"id"   gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"unique;not null"`
}

type Category struct {
	ID   uint   `json:"id"   gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"unique;not null"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}

func NewValidator() *validator.Validate {
	return validator.New()
}

