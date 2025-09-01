package entity

import (
	"gorm.io/gorm"
)

type ProductPostCategory struct {
	gorm.Model

	Name 		string 			`json:"name"`
}
