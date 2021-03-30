package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-template/app/models"
	"golang-template/app/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (app *App) Login(c *gin.Context) {
	var user models.User
	if err := app.DB.Where("email = ?", c.Param("email")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not found!"})
		return
	}
	//Compare the password form and database if match
	bytePassword := []byte(user.Password)
	byteHashedPassword := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Password"})

	}
	//Generate the JWT auth token
	token,err := utils.CreateToken(uint32(user.ID))
	fmt.Println(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot create token"})

	}
	return
}
func (app *App) GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome:)",
	})
	return
}
func (app *App) NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Not Found",
	})
	return
}

func (app *App) CreateCar(c *gin.Context) {
	// Validate input
	var car models.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	app.DB.Create(&car)
	c.JSON(http.StatusOK, gin.H{"car": car})
}
func (app *App) UpdateCar(c *gin.Context) {
	var car models.Car
	if err := app.DB.Where("id = ?", c.Param("id")).First(&car).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var updateCar models.Car
	if err := c.ShouldBindJSON(&updateCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	app.DB.Model(&car).Updates(updateCar)
	c.JSON(http.StatusOK, gin.H{"car": car})
}
func (app *App) DeleteBook(c *gin.Context) {
	// Get model if exist
	var car models.Car
	if err := app.DB.Where("id = ?", c.Param("id")).First(&car).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	app.DB.Delete(&car)
	c.JSON(http.StatusOK, gin.H{"delete car": true})
}
func (app *App) FindCars(c *gin.Context) {
	var cars []models.Car
	app.DB.Find(&cars)
	c.JSON(http.StatusOK, gin.H{"cars": cars})

}
func (app *App) FindCar(c *gin.Context) {
	// Get model if exists
	var car models.Car
	if err := app.DB.Where("id = ?", c.Param("id")).First(&car).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": car})

}
