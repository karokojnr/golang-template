package models

import (
	"github.com/jinzhu/gorm"
)
type Car struct {
	gorm.Model
	RegNo        string `gorm:"unique; not null" json:"reg_no" binding:"required"`
	VehicleModel string `gorm:"not null" json:"vehicle_model" binding:"required"`
}
//gorm:"default:'KAA100A
//gorm:"type:varchar(100);not null"
