package entity

import("gorm.io/gorm")

type Person struct{
	gorm.Model

	Fristname		string		`json:"fristname"`

	Lastname		string		`json:"lastname"`

	Email			string		`gorm:"uniqueIndex" json:"email"`

	Age				int			`josn:"age"`

	Phone			string		`json:"phone"`

	Birthday		string		`json:"birthday"`

	Address			string		`json:"address"`

	GenderID		uint
	Gender			*Gender

	MemberID		uint
	Member			*Member
}