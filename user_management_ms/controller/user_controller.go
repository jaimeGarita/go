package controller

import (
	"net/http"
	"user_management_ms/model"
	"user_management_ms/service"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	usrs := service.GetUsersService()

	c.JSON(http.StatusOK, gin.H{"users": usrs})
}

func SaveUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	usrDB, err := service.SaveUser(&user)
	if err == nil {
		c.JSON(http.StatusOK, &usrDB)
		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

}
