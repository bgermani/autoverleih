package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bgermani/autoverleih/model"

	"github.com/bgermani/autoverleih/config"
)

type CreateAutoInput struct {
	Brand string `json:"brand" binding:"required"`
	Model string `json:"model" binding:"required"`
}

type UpdateAutoInput struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
}

// GET /auto
func FindAutos(c *gin.Context) {
	var autos []model.Auto
	config.DB.Find(&autos)

	c.JSON(http.StatusOK, gin.H{"data": autos})
}

// GET /auto/:id
func FindAuto(c *gin.Context) { // Get model if exist
	var auto model.Auto

	if err := config.DB.Where("id = ?", c.Param("id")).First(&auto).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auto not found."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": auto})
}

// POST /auto
func CreateAuto(c *gin.Context) {
	var input CreateAutoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	auto := model.Auto{Brand: input.Brand, Model: input.Model}
	config.DB.Create(&auto)

	c.JSON(http.StatusOK, gin.H{"data": auto})
}

// PATCH /auto/:id
func UpdateAuto(c *gin.Context) {
	var auto model.Auto
	if err := config.DB.Where("id = ?", c.Param("id")).First(&auto).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auto not found."})
		return
	}

	var input UpdateAutoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&auto).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": auto})
}

// DELETE /auto/:id
func DeleteAuto(c *gin.Context) {
	var auto model.Auto
	if err := config.DB.Where("id = ?", c.Param("id")).First(&auto).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auto not found."})
		return
	}

	config.DB.Delete(&auto)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
