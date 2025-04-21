package database

import (
	"api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	conn, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/laravel"), &gorm.Config{})
	if err != nil {
		panic("could not connect to database")
	}

	DB = conn

	conn.AutoMigrate(
		&models.Admin{}, &models.Category{}, &models.Product{}, &models.User{}, &models.ProductDetail{}, &models.Review{},&models.Favorite{}, &models.Cart{}, &models.Order{}, &models.BuktiPembayaran{})
}
