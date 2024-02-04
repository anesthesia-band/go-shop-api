package models

import (
	"gorm.io/gorm"
)

type ItemType struct {
	gorm.Model
	Id    uint   `gorm:"primaryKey"`
	Title string `gorm:"index"`
	Name  string
}
