package entity

import (
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model

	Description       	string  		`json:"description"`

	ReportTypeID   		uint  			`json:"report_type_id"`
	ReportType  		*ReportType  	`gorm:"foreignKey:ReportTypeID" json:"report_type"`

	MemberID    		uint    		`json:"member_id"`
	Member       		*Member 		`gorm:"foreignKey:MemberID" json:"member"`

	SellerID    		uint    		`json:"seller_id"`
	Seller       		*Seller 		`gorm:"foreignKey:SellerID" json:"seller"`

	PaymentID    		uint    		`json:"payment_id"`
	Payment     		*Payment 		`gorm:"foreignKey:PaymentID" json:"payment"`
}
