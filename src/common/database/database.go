package database

import (
	"go-shop-api/src/repository/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	db, err := gorm.Open(postgres.Open("host=db user=root password=root dbname=shop_db port=5432"), &gorm.Config{})

	if err != nil {
		panic("db not connected")
	}

	autoMigrate(db)

	return db
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(models.ItemType{}, models.Item{}, models.GroupType{}, models.Group{}, models.GroupItem{}, models.Order{})
}
