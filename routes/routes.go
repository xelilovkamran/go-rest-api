package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	authorized := server.Group("/")
	authorized.Use(middlewares.AuthMiddleware())
	{
		authorized.GET("/v1/events", getEvents)
		authorized.POST("/v1/events", createEvent)
		authorized.GET("/v1/events/:id", getEventByID)
		authorized.PUT("/v1/events/:id", updateEvent)
		authorized.DELETE("/v1/events/:id", deleteEvent)

		authorized.POST("/v1/events/:id/register", registerForEvent)
		authorized.DELETE("/v1/events/:id/register", cancelRegistration)
	}

	server.POST("/v1/signup", signup)
	server.POST("/v1/login", login)
}