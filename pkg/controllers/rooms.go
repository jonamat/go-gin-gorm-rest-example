package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonamat/go-gin-gorm-rest-example/pkg/db"
	"github.com/jonamat/go-gin-gorm-rest-example/pkg/models"
)

// GET /rooms
func FindRooms(c *gin.Context) {
	var rooms []models.Room

	// Get records
	if err := db.Gorm.Find(&rooms).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while fetching data"})
	}

	c.JSON(http.StatusOK, gin.H{"data": rooms})
}

// GET /rooms/:id
func FindRoom(c *gin.Context) {
	var room models.Room

	// Get record
	if err := db.Gorm.First(&room, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": room})
}

// POST /rooms
func CreateRoom(c *gin.Context) {
	var input models.Room

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

// PUT /rooms/:id
func UpdateRoom(c *gin.Context) {
	var room models.Room
	var input models.Room

	// Check if record exist
	if err := db.Gorm.First(&room, c.Param("id")).Error; err != nil {
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
	if err := db.Gorm.Model(&room).Updates(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": room})
}

// DELETE /rooms/:id
func DeleteRoom(c *gin.Context) {
	var room models.Room

	// Get record if exist
	if err := db.Gorm.First(&room, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	// Delete record
	if err := db.Gorm.Delete(&room).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
