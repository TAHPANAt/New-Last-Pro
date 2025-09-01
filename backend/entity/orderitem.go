package entity

import (
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model

	OrderID			uint		`json:"order_id"`
	Order			*Order		`gorm:"foreignKey:OrderID" json:"order"`

	ProductID		uint		`json:"product_id"`
	Product			*Product	`gorm:"foreignKey:ProductID" json:"product"`	

	Quantity		int			`json:"quantity"`

	PriceAtOrder	float64		`json:"price_at_order"`
}
