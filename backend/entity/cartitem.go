package entity

import (
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model

	CartID		uint		`json:"cart_id"`
	Cart		*Cart		`gorm:"foreignKey:CartID" json:"cart"`

	ProductID	uint		`json:"product_id"`
	Product		*Product	`gorm:"foreignKey:ProductID" json:"product"`

	Quantity	int			`json:"quantity"`
}	