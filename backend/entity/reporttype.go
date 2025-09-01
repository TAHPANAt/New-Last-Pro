package entity

import (
	"gorm.io/gorm"
)

type ReportType struct {
	gorm.Model

	Name        string  		`json:"name"`

}
