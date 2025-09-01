package configs

import (
//	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"project.com/sa-68-project/entity"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("sa-example.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = database
}

func SetupDatabase() {
	// Migrate the schema
	db.AutoMigrate(
		&entity.ActivityLog{},
		&entity.Admin{},
		&entity.Cart{},
		&entity.CartItem{},
		&entity.CreateCode{},
		&entity.DiscountCode{},
		&entity.Gender{},
		&entity.Member{},
		&entity.MemberProfile{},
		&entity.Order{},
		&entity.OrderDiscount{},
		&entity.OrderItem{},
		&entity.Payment{},
		&entity.PaymentMethod{},
		&entity.Person{},
		&entity.PostANewProduct{},
		&entity.Product{},
		&entity.ProductDiscount{},
		&entity.ProductImage{},
		&entity.ProductPostCategory{},
		&entity.Report{},
		&entity.ReportType{},
		&entity.Seller{},
		&entity.ShopAddress{},
		&entity.ShopCategory{},
		&entity.ShopProfile{},
	)
}