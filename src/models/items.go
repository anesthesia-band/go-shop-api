package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Id         uint `gorm:"primaryKey"`
	Name       string
	ItemTypeId uint `gorm:"index"`
	Data       string
	Price      string
	Active     bool `gorm:"default:true"`
	ItemType
}
