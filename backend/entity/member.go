package entity

import	(
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model

	Username		string		`gorm:"uniqueIndex" json:"username"`

	Password		string		`json:"password"`

	People			Person

	Seller			Seller

	Cart			Cart
}