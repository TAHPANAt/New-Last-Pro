package entity
import ("gorm.io/gorm")

type Gender struct{
	gorm.Model

	Gender		string		`json:"gender"`

	People		[]Person	`gorm:"foreignKey:GenderID" json:"person"`
}
