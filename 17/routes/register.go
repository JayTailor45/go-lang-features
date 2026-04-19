package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"jaytailor.com/event-management/models"
)

func registerForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch the event"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user to event"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User successfully registered"})
}

func cancelRegistration(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch the event"})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel user registration"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User registration successfully cancelled"})
}
