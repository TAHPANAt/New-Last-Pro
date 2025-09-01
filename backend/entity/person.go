package entity

import("gorm.io/gorm")

type Person struct{
	gorm.Model

	Firstname		string		`json:"firstname"`

	Lastname		string		`json:"lastname"`

	Email			string		`gorm:"uniqueIndex" json:"email"`

	Age				int			`json:"age"`

	Phone			string		`json:"phone"`

	Birthday		string		`json:"birthday"`

	Address			string		`json:"address"`

	GenderID		uint		`json:"gender_id"`
	Gender			*Gender		`gorm:"foreignKey:GenderID" json:"gender"`

}