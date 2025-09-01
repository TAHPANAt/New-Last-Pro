package entity

import (
	"gorm.io/gorm"
)

type ShopCategory struct {
	gorm.Model

	Name        string  		`json:"name"`

}
