package entity

import (
	"gorm.io/gorm"
)

type ActivityLog struct {
	gorm.Model

	MemberID uint   `json:"member_id"`
	Member   *Member `gorm:"foreignKey:MemberID" json:"member"`

	Action    string `json:"action"`
	Timestamp string `json:"timestamp"`
}
