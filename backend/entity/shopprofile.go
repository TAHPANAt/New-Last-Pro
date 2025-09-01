package entity

import (
	"gorm.io/gorm"
)

type ShopProfile struct {
	gorm.Model

	ShopName        string       `json:"shop_name"`
	
	ShopDescription string       `json:"shop_description"`

	Logo             string       `json:"logo"`

	SellerID        uint         `json:"seller_id"`
	Seller          *Seller      `gorm:"foreignKey:SellerID" json:"seller"`

	ShopCategoryID  uint         `json:"shop_category_id"`
	ShopCategory    *ShopCategory `gorm:"foreignKey:ShopCategoryID" json:"shop_category"`

}
