package entity

import(
	"gorm.io/gorm"
)

type Seller struct {
	gorm.Model

	Shopname		string		`json:"shopname"`

	Shopaddress		string		`json:"shopaddress"`

	MemberID		uint
	Member			*Member

	ProductID		uint
	Product			*Product			
}