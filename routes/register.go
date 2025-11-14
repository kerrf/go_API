package routes

import (
	"net/http"
	"strconv"

	"example.com/RestAPI/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id.", "error": err.Error()})
		return
	}

	event, err := models.GetEvent(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	err = event.Register(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered!"})
}

func cancelRegistration(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id.", "error": err.Error()})
		return
	}

	var event models.Event
	event.ID = eventID

	err = event.CancelRegistration(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered!"})
}
