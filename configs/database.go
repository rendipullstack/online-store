package configs

import (
	"fmt"

	"online-store/helpers"
	"online-store/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDatabase(username, password, host, port, dbName string) (*gorm.DB, error) {
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})

	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Category{}, &models.Cart{}, &models.TransactionHistory{})

	return db, helpers.ReturnIfError(err)
}
