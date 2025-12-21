package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func registerForEvent(c *gin.Context) {
	userId := int(c.Keys["claims"].(jwt.MapClaims)["userId"].(float64))

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID."})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found."})
		return
	}

	err = event.Register(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for event."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully registered for event."})
}

func cancelRegistration(c *gin.Context) {
	userId := int(c.Keys["claims"].(jwt.MapClaims)["userId"].(float64))

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID."})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found."})
		return
	}

	err = event.CancelRegistration(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully canceled registration."})
}