package controllers

import "gorm.io/gorm"

type App struct {
	DB *gorm.DB
}
