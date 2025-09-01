package entity

import (
	"gorm.io/gorm"
)

type DiscountCode struct {
	gorm.Model

	Percentage  	float64 	`json:"percentage"`

	ExpiryDate  	string  	`json:"expiry_date"`

	UsageCount  	int     	`json:"usage_count"`

	Products   		[]Product 	`gorm:"many2many:product_discounts" json:"products"`

	CreateCodeID	uint		`json:"create_code_id"`
	CreateCode		*CreateCode	`gorm:"foreignKey:CreateCodeID" json:"create_code"`
}
