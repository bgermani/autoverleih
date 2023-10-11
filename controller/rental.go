package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bgermani/autoverleih/model"

	"github.com/bgermani/autoverleih/config"

	"time"

	"gorm.io/datatypes"
)

type CreateRentalInput struct {
	AutoId         int    `json:"auto_id" binding:"required"`
	CustomerId     int    `json:"customer_id" binding:"required"`
	KilometerCount int    `json:"kilometer_count" binding:"required"`
	Start          string `json:"period_start" binding:"required"`
	End            string `json:"period_end" binding:"required"`
}

type UpdateRentalInput struct {
	AutoId         int            `json:"auto_id"`
	CustomerId     int            `json:"customer_id"`
	KilometerCount int            `json:"kilometer_count"`
	Start          datatypes.Date `json:"period_start"`
	End            datatypes.Date `json:"period_end"`
}

// GET /rental
func FindRentals(c *gin.Context) {
	var rentals []model.Rental
	config.DB.Find(&rentals)

	c.JSON(http.StatusOK, gin.H{"data": rentals})
}

// GET /rental/:id
func FindRental(c *gin.Context) {
	var rental model.Rental

	if err := config.DB.Where("id = ?", c.Param("id")).First(&rental).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Rental not found."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rental})
}

// GET /rental/active
func FindActiveRentals(c *gin.Context) {
	var active []model.Rental

	config.DB.Where("period_start <= ? AND period_end >= ?", time.Now(), time.Now()).Find(&active)

	c.JSON(http.StatusOK, gin.H{"data": active})
}

// GET /rental/active-count
func FindActiveRentalCount(c *gin.Context) {
	var active []model.Rental

	count := config.DB.Where("period_start <= ? AND period_end >= ?", time.Now(), time.Now()).Find(&active)

	c.JSON(http.StatusOK, gin.H{"data": count.RowsAffected})
}

// POST /rental
func CreateRental(c *gin.Context) {
	var input CreateRentalInput

	const layout = "2006-01-02"
	start, _ := time.Parse(layout, input.Start)
	end, _ := time.Parse(layout, input.End)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO STILL NEED LOGIC TO MAKE SURE CAR ISNT ALREADY RENTED

	rental := model.Rental{
		AutoId:         input.AutoId,
		CustomerId:     input.CustomerId,
		KilometerCount: input.KilometerCount,
		Start:          datatypes.Date(start),
		End:            datatypes.Date(end),
	}
	config.DB.Create(&rental)

	c.JSON(http.StatusOK, gin.H{"data": rental})
}

// PATCH /rental/:id
func UpdateRental(c *gin.Context) {
	var rental model.Rental
	if err := config.DB.Where("id = ?", c.Param("id")).First(&rental).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Rental not found."})
		return
	}

	// TODO STILL NEED LOGIC TO MAKE SURE CAR ISNT ALREADY RENTED

	var input UpdateRentalInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&rental).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": rental})
}

// DELETE /rental/:id
func DeleteRental(c *gin.Context) {
	var rental model.Rental
	if err := config.DB.Where("id = ?", c.Param("id")).First(&rental).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Rental not found."})
		return
	}

	config.DB.Delete(&rental)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
