package entity

import (
	"gorm.io/gorm"
)

type ProductImage struct {
	gorm.Model

	ProductID uint   `json:"product_id"`
	Product   *Product `gorm:"foreignKey:ProductID" json:"product"`
	
	ImageURL  string `json:"image_url"`
}
