package routes

import (
	"user_management_ms/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoles(r *gin.Engine) {
	r.GET("/users", controller.GetUsers)
}
