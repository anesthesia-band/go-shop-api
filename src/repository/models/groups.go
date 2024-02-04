package models

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Id          uint `gorm:"primaryKey"`
	Name        string
	GroupTypeId uint `gorm:"index"`
	Data        string
	Price       string
	Active      bool `gorm:"default:true"`
	GroupType
}
