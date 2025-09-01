package entity

import (
	"gorm.io/gorm"
)

type MemberProfile struct {
	gorm.Model

	MemberID   uint   `json:"member_id"`
	Member     *Member `gorm:"foreignKey:MemberID" json:"member"`

	AvatarURL  string `json:"avatar_url"`
}
