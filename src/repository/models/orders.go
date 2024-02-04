package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Id        uint `gorm:"primaryKey"`
	Name      string
	Phone     sql.NullString `gorm:"index"`
	Email     string         `gorm:"index"`
	OrderData string
	Status    string `gorm:"index;default:new"`
}
