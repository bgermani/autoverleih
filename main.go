package main

import (
	"github.com/gin-gonic/gin"

	"github.com/bgermani/autoverleih/config"

	"github.com/bgermani/autoverleih/controller"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()

	r.GET("/auto", controller.FindAutos)
	r.GET("/auto/:id", controller.FindAuto)
	r.POST("/auto", controller.CreateAuto)
	r.PATCH("/auto/:id", controller.UpdateAuto)
	r.DELETE("/auto/:id", controller.DeleteAuto)

	r.GET("/customer", controller.FindCustomers)
	r.GET("/customer/:id", controller.FindCustomer)
	r.POST("/customer", controller.CreateCustomer)
	r.PATCH("/customer/:id", controller.UpdateCustomer)
	r.DELETE("/customer/:id", controller.DeleteCustomer)

	r.GET("/rental", controller.FindRentals)
	r.GET("/rental/:id", controller.FindRental)
	r.GET("/rental/active", controller.FindActiveRentals)
	r.GET("/rental/active-count", controller.FindActiveRentalCount)
	r.POST("/rental", controller.CreateRental)
	r.PATCH("/rental/:id", controller.UpdateRental)
	r.DELETE("/rental/:id", controller.DeleteRental)

	r.Run()
}
