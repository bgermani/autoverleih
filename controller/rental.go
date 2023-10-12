package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bgermani/autoverleih/model"

	"github.com/bgermani/autoverleih/config"

	"time"
)

type CreateRentalInput struct {
	AutoId         int    `json:"auto_id" binding:"required"`
	CustomerId     int    `json:"customer_id" binding:"required"`
	KilometerCount int    `json:"kilometer_count" binding:"required"`
	Start          string `json:"period_start" binding:"required"`
	End            string `json:"period_end" binding:"required"`
}

type UpdateRentalInput struct {
	AutoId         int    `json:"auto_id"`
	CustomerId     int    `json:"customer_id"`
	KilometerCount int    `json:"kilometer_count"`
	Start          string `json:"period_start"`
	End            string `json:"period_end"`
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

	config.DB.Where("start <= ? AND end >= ?", time.Now(), time.Now()).Find(&active)

	c.JSON(http.StatusOK, gin.H{"data": active})
}

// GET /rental/active-count
func FindActiveRentalCount(c *gin.Context) {
	var active []model.Rental

	count := config.DB.Where("start <= ? AND end >= ?", time.Now(), time.Now()).Find(&active)

	c.JSON(http.StatusOK, gin.H{"data": count.RowsAffected})
}

// POST /rental
func CreateRental(c *gin.Context) {
	var input CreateRentalInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	const layout = "2006-01-02 15:04:05"
	start, _ := time.Parse(layout, input.Start)
	end, _ := time.Parse(layout, input.End)

	var auto model.Auto
	if err := config.DB.Where("id = ?", input.AutoId).First(&auto).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auto not found."})
		return
	}

	var customer model.Customer
	if err := config.DB.Where("id = ?", input.CustomerId).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer not found."})
		return
	}

	var activeRental model.Rental
	config.DB.Where("auto_id = ? AND (start BETWEEN ? AND ? OR end BETWEEN ? AND ?)", input.AutoId, start, end, start, end).Limit(1).Find(&activeRental)
	if activeRental.AutoId == input.AutoId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auto already rented during selected period."})
		return
	}

	rental := model.Rental{
		AutoId:         input.AutoId,
		CustomerId:     input.CustomerId,
		KilometerCount: input.KilometerCount,
		Start:          start,
		End:            end,
		ModifiedAt:     time.Now(),
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

	// UNFINISHED

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
