package services

import (
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(connection *gorm.DB) {
	db = connection
}
