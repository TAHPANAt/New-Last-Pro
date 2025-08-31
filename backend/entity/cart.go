package entity

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model

	MemberID		uint
	Member			*Member
}