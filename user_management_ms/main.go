package main

import (
	"log"
	"user_management_ms/config"
	"user_management_ms/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	routes.SetupRoles(r)

	log.Println("🚀 Servidor corriendo en http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error iniciando servidor:", err)
	}

}
