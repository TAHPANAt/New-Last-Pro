package entity

import (
	"gorm.io/gorm"
)

type CreateCode struct {
	gorm.Model

	Codenamee        string  	`json:"codenamee"`

	Percentage  	float64 	`json:"percentage"`

	ExpiryDate  	string 		`json:"expiry_date"`

	UsageCount  	int     	`json:"usage_count"`

	SellerID		uint		`json:"seller_id"`
	Seller			*Seller		`gorm:"foreignKey:SellerID" json:"seller"`
}
