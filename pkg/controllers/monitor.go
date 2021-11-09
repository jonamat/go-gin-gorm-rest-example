package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST /ping
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
