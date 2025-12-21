package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve events."})
		return
	}
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var event models.Event
	if err := c.BindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event data."})
		return
	}

	userId, ok := c.Keys["claims"].(jwt.MapClaims)["userId"].(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve user ID from token."})
		return
	}

	event.UserID = int(userId)

	if err := event.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event."})
		return
	}
	c.JSON(http.StatusCreated, event)
}

func getEventByID(c *gin.Context) {
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
	c.JSON(http.StatusOK, event)
}

func updateEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID."})
		return
	}

	userId, ok := c.Keys["claims"].(jwt.MapClaims)["userId"].(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve user ID from token."})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found."})
		return
	}

	if event.UserID != int(userId) {
		c.JSON(http.StatusForbidden, gin.H{"message": "You do not have permission to update this event."})
		return
	}

	var updatedEvent models.Event
	if err := c.BindJSON(&updatedEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event data."})
		return
	}
	updatedEvent.ID = id
	updatedEvent.UserID = int(userId)

	if err := updatedEvent.UpdateEvent(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event."})
		return
	}
	c.JSON(http.StatusOK, updatedEvent)
}

func deleteEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID."})
		return
	}

	userId, ok := c.Keys["claims"].(jwt.MapClaims)["userId"].(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve user ID from token."})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found."})
		return
	}
	if event.UserID != int(userId) {
		c.JSON(http.StatusForbidden, gin.H{"message": "You do not have permission to delete this event."})
		return
	}

	err = event.DeleteEvent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully."})
}