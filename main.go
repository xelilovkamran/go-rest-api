package main

import (
	"github.com/gin-gonic/gin"

	"example.com/rest-api/db"
	"example.com/rest-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}

