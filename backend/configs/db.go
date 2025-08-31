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
		&entity.Cart{},
		&entity.Gender{},
		&entity.Member{},
		&entity.Person{},
		&entity.Product{},
		&entity.Seller{},
	)
}