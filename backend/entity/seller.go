package entity

import(
	"gorm.io/gorm"
)

type Seller struct {
	gorm.Model

	MemberID		uint			`json:"member_id"`
	Member			*Member			`gorm:"foreignKey:MemberID" json:"member"`

	ShopAddressID	uint			`json:"shop_address_id"`
	ShopAddress		*ShopAddress	`gorm:"foreignKey:ShopAddressID" json:"shop_address"`

}