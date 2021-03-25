package database

import (
	"fmt"
	"golang-template/app/models"
	"golang-template/app/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=karokojnr dbname=cars port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database!")
	}
	return db, nil
}
func AutoMigrateDB(db *gorm.DB) {
	utils.Log("auto-migrations running...")
	db.AutoMigrate(&models.Car{})
	utils.Log("auto-migration complete...")
}
