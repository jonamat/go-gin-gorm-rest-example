package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonamat/go-gin-gorm-rest-example/pkg/db"
	"github.com/jonamat/go-gin-gorm-rest-example/pkg/models"
)

// GET /bookings
func FindBookings(c *gin.Context) {
	var bookings []models.Booking

	// Get records
	if err := db.Gorm.Preload("Services").Find(&bookings).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while fetching data"})
		return
	}

	// Populate service_ids
	for index := range bookings {
		for _, service := range bookings[index].Services {
			bookings[index].ServiceIDs = append(bookings[index].ServiceIDs, service.ID)
		}

		// If service_ids is nil, set it to empty array
		if len(bookings[index].ServiceIDs) == 0 {
			bookings[index].ServiceIDs = make([]uint, 0)
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": bookings})
}

// GET /bookings/:id
func FindBooking(c *gin.Context) {
	var booking models.Booking

	// Get record
	if err := db.Gorm.First(&booking, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	// Populate service_ids
	for _, service := range booking.Services {
		booking.ServiceIDs = append(booking.ServiceIDs, service.ID)
	}

	// If service_ids is nil, set it to empty array
	if len(booking.ServiceIDs) == 0 {
		booking.ServiceIDs = make([]uint, 0)
	}

	c.JSON(http.StatusOK, gin.H{"data": booking})
}

// POST /bookings
func CreateBooking(c *gin.Context) {
	var input models.Booking

	// Validate schema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate input
	if err := input.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if all provided services are available and populate Services
	for _, serviceID := range input.ServiceIDs {
		var service models.Service
		if err := db.Gorm.First(&service, serviceID).Error; err != nil {
			log.Println(err)
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("service with id %d not found", serviceID)})
			return
		}
		input.Services = append(input.Services, service)
	}

	// Create record
	if err := db.Gorm.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// PUT /services/:id
func UpdateBooking(c *gin.Context) {
	var booking models.Booking
	var input models.Booking

	// Check if record exist
	if err := db.Gorm.First(&booking, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	// Validate schema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate input
	if err := input.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO delegate to create
	// Check if all provided services are available and populate Services
	for _, serviceID := range input.ServiceIDs {
		var service models.Service
		if err := db.Gorm.First(&service, serviceID).Error; err != nil {
			log.Println(err)
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("service with id %d not found", serviceID)})
			return
		}
		input.Services = append(input.Services, service)
	}

	// Update record
	if err := db.Gorm.Model(&booking).Updates(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// DELETE /bookings/:id
func DeleteBooking(c *gin.Context) {
	var booking models.Booking

	// Get record if exist
	if err := db.Gorm.Where("id = ?", c.Param("id")).First(&booking).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	// Delete record
	if err := db.Gorm.Delete(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
