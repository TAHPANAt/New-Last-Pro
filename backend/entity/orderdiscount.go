package entity

import (
	"gorm.io/gorm"
)

type OrderDiscount struct {
	gorm.Model

	OrderID       	uint      		`json:"order_id"`
	Order       	*Order    		`gorm:"foreignKey:OrderID" json:"order"`

	DiscountCodeID 	uint      		`json:"discount_code_id"`
	DiscountCode   	*DiscountCode 	`gorm:"foreignKey:DiscountCodeID" json:"discount_code"`
}
