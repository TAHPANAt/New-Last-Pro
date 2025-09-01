package entity

import (
	"gorm.io/gorm"
)

type ShopAddress struct {
	gorm.Model

	Province    string `json:"province"`

	City       string `json:"city"`

	AddressLine string `json:"address_line"`

	PostalCode  string `json:"postal_code"`

	Country    string `json:"country"`

	District   string `json:"district"`
}
