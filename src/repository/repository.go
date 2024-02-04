package repository

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(connection *gorm.DB) {
	DB = connection
}
