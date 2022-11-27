package config

import (
	"go-cron/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func StartDB() {
	var err error
	dsn := "root:root@tcp(localhost)/assignment-3?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.Element{})

}

func GetDB() *gorm.DB {
	return DB
}
