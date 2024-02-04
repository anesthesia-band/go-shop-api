package models

import (
	"gorm.io/gorm"
)

type GroupItem struct {
	gorm.Model
	GroupId  uint `gorm:"primaryKey"`
	ItemId   uint `gorm:"primaryKey"`
	Quantity uint `gorm:"default:0"`
	Active   bool `gorm:"default:true"`
	Group
	Item
}
