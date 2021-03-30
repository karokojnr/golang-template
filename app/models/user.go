package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique; not null" json:"email" binding:"required"`
	Password string `gorm:"not null" json:"password" binding:"required"`
}

