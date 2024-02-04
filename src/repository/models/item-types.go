package models

import (
	"gorm.io/gorm"
)

type ItemType struct {
	gorm.Model
	Id   uint `gorm:"primaryKey"`
	Name string
}
