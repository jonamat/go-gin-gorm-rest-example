package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonamat/go-gin-gorm-rest-example/pkg/db"
	"github.com/jonamat/go-gin-gorm-rest-example/pkg/models"
)

// GET /services
func FindServices(c *gin.Context) {
	var services []models.Service

	// Get records
	if err := db.Gorm.Find(&services).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while fetching data"})
	}

	c.JSON(http.StatusOK, gin.H{"data": services})
}

// GET /services/:id
func FindService(c *gin.Context) {
	var service models.Service

	// Get record
	if err := db.Gorm.First(&service, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": service})
}

// POST /services
func CreateService(c *gin.Context) {
	var input models.Service

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

	// Create record
	if err := db.Gorm.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// PUT /services/:id
func UpdateService(c *gin.Context) {
	var service models.Service
	var input models.Service

	// Check if record exist
	if err := db.Gorm.First(&service, c.Param("id")).Error; err != nil {
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

	// Update record
	if err := db.Gorm.Model(&service).Updates(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": service})
}

// DELETE /services/:id
func DeleteService(c *gin.Context) {
	var service models.Service

	// Get record if exist
	if err := db.Gorm.First(&service, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	// Delete record
	if err := db.Gorm.Delete(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
