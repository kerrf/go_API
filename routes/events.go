package routes

import (
	"net/http"
	"strconv"

	"example.com/RestAPI/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id.", "error": err.Error()})
		return
	}
	event, err := models.GetEvent(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse input.", "error": err.Error()})
		return
	}

	event.UserID = context.GetInt64("userID")

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id.", "error": err.Error()})
		return
	}

	userID := context.GetInt64("UserID")
	event, err := models.GetEvent(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events.", "error": err.Error()})
	}

	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event."})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindBodyWithJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse input.", "error": err.Error()})
		return
	}

	updatedEvent.ID = eventID
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully."})

}

func deleteEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id.", "error": err.Error()})
		return
	}

	event, err := models.GetEvent(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events.", "error": err.Error()})
		return
	}

	userID := context.GetInt64("UserID")
	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event."})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event successfully deleted."})
}
