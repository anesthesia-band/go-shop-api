package models

import (
	"gorm.io/gorm"
)

type GroupType struct {
	gorm.Model
	Id   uint `gorm:"primaryKey"`
	Name string
}
