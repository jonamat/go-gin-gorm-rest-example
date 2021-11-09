package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jonamat/go-gin-gorm-rest-example/pkg/db"
	"github.com/jonamat/go-gin-gorm-rest-example/pkg/models"
)

// GET /users/:id
func FindUser(c *gin.Context) {
	var user models.User

	session := sessions.Default(c)
	userID := session.Get("id")

	if fmt.Sprintf("%d", userID) != c.Param("id") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// Get record
	if err := db.Gorm.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// PUT /users/:id
func UpdateUser(c *gin.Context) {
	var user models.User
	var input models.User

	// Check if record exist
	if err := db.Gorm.First(&user, c.Param("id")).Error; err != nil {
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
	if err := db.Gorm.Model(&user).Updates(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
