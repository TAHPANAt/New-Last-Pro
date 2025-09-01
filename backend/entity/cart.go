package entity

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model

	MemberID   		uint      		`json:"member_id"`
	Member     		*Member   		`gorm:"foreignKey:MemberID" json:"member"`

	CartItems  		[]CartItem 		`gorm:"foreignKey:CartID" json:"cart_items"`
}
