package entity

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model

	Username 		string 		`json:"username"`

	Password 		string 		`json:"password"`

	PersonID 		uint   		`json:"person_id"`
	Person   		*Person 	`gorm:"foreignKey:PersonID" json:"person"`
}
