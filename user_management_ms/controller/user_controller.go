package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := []map[string]string{
		{"id": "1", "name": "alice"},
		{"id": "2", "name": "Bob"},
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}
