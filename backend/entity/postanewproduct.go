package entity

import (
	"gorm.io/gorm"
)

type PostANewProduct struct {
	gorm.Model

	ProductPostCategoryID 	uint              			`json:"product_post_category_id"`
	ProductPostCategory    	*ProductPostCategory 		`gorm:"foreignKey:ProductPostCategoryID" json:"product_post_category"`

	ProductID 				uint        				`json:"product_id"`
	Product   				*Product    				`gorm:"foreignKey:ProductID" json:"product"`
}
