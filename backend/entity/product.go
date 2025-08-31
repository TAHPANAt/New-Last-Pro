package entity

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	Productname			string		`json:"productname"`

	Productdescription	string		`json:"productdescription"`

	Price				int			`json:"price"`

	Quantity			int			`json:"quantity"`

	Sellers				[]Seller
}