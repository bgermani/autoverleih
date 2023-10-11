package config

import (
	"github.com/bgermani/autoverleih/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := "user:password@tcp(db:3306)/db?charset=utf8&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&model.Auto{})
	if err != nil {
		return
	}

	err = database.AutoMigrate(&model.Customer{})
	if err != nil {
		return
	}

	err = database.AutoMigrate(&model.Rental{})
	if err != nil {
		return
	}

	DB = database
}
