package entity

import (
	"gorm.io/gorm"
)

type ProductDiscount struct {
	gorm.Model

	ProductID     	uint    		`json:"product_id"`
	Product       	*Product 		`gorm:"foreignKey:ProductID" json:"product"`

	DiscountCodeID 	uint    		`json:"discount_code_id"`
	DiscountCode   	*DiscountCode 	`gorm:"foreignKey:DiscountCodeID" json:"discount_code"`
}
