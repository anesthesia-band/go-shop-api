package models

import (
	"gorm.io/gorm"
)

type GroupItem struct {
	gorm.Model
	GroupId  uint `gorm:"primaryKey"`
	ItemId   uint `gorm:"primaryKey"`
	quantity uint `gorm:"default:0"`
	Group
	Item
}
