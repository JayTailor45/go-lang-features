package main

import (
	"github.com/gin-gonic/gin"
	"jaytailor.com/event-management/db"
	"jaytailor.com/event-management/routes"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
