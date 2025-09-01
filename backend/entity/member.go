package entity

import	(
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model

	Username		string		`gorm:"uniqueIndex" json:"username"`

	Password		string		`json:"password"`

	PersonID		uint		`json:"person_id"`
	Person			*Person		`gorm:"foreignKey:PersonID" json:"person"`  

}