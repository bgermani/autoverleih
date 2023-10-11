package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bgermani/autoverleih/model"

	"github.com/bgermani/autoverleih/config"
)

type CreateCustomerInput struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

type UpdateCustomerInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// GET /customer
func FindCustomers(c *gin.Context) {
	var customers []model.Customer
	config.DB.Find(&customers)

	c.JSON(http.StatusOK, gin.H{"data": customers})
}

// GET /customer/:id
func FindCustomer(c *gin.Context) {
	var customer model.Customer

	if err := config.DB.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer not found."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// POST /customer
func CreateCustomer(c *gin.Context) {
	var input CreateCustomerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer := model.Customer{FirstName: input.FirstName, LastName: input.LastName}
	config.DB.Create(&customer)

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// PATCH /customer/:id
func UpdateCustomer(c *gin.Context) {
	var customer model.Customer
	if err := config.DB.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer not found."})
		return
	}

	var input UpdateCustomerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&customer).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// DELETE /customer/:id
func DeleteCustomer(c *gin.Context) {
	var customer model.Customer
	if err := config.DB.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer not found."})
		return
	}

	config.DB.Delete(&customer)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
